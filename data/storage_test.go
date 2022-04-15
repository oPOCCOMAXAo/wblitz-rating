package data

import (
	"github.com/stretchr/testify/suite"
)

type StorageSuite struct {
	Storage
	suite.Suite
}

func (s *StorageSuite) TestPlayer() {
	data := []Player{
		{ID: 1, Nickname: "test", Battles: 123, Damage: 12345, Wins: 67},
		{ID: 5, Nickname: "test2", Battles: 122, Damage: 12344, Wins: 66},
	}

	err := s.Storage.SavePlayers(data)
	s.Require().NoError(err)

	loaded, err := s.Storage.LoadPlayers(&BatchConfig{
		Limit: 200,
	})
	s.Require().NoError(err)
	s.Require().Equal(data, loaded)
}

func (s *StorageSuite) TestRating() {
	data := []Rating{
		{Number: 1, SpaID: 5, Score: 3, MMR: 1.05},
		{Number: 2, SpaID: 1, Score: 2, MMR: 1.01},
	}

	err := s.Storage.SaveRating(data)
	s.Require().NoError(err)

	loaded, err := s.Storage.LoadRating(&BatchConfig{
		Limit: 200,
	})
	s.Require().NoError(err)
	s.Require().Equal(data, loaded)
}
