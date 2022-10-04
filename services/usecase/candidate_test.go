package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/services/repository/postgres/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_Candidate_UpdateCandidate(t *testing.T) {
	type testCase struct {
		desc                 string
		wantError            bool
		input                models.CandidateArgument
		onCheckCandidateById func(mock *mocks.MockCandidateRepository)
		onUpdateCandidate    func(mock *mocks.MockCandidateRepository)
		onGetDetailCandidate func(mock *mocks.MockCandidateRepository)
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		desc:      "failed, server error",
		wantError: true,
		input: models.CandidateArgument{
			ID:               "xxx",
			FirstName:        "dhany",
			LastName:         "aritonang",
			PhoneNumber:      "080821",
			TelpNumber:       "080821",
			Address:          "simanabun",
			ProfilPictureUrl: "gambar.jpg",
		},
		onCheckCandidateById: func(mock *mocks.MockCandidateRepository) {
			mock.EXPECT().CheckCandidateById(gomock.Any(), gomock.Any()).Return(
				models.Candidate{},
				errors.New("internal server error"),
			)
		},
	})

	testCases = append(testCases, testCase{
		desc:      "failed, internal serverr error when update candidate",
		wantError: true,
		input: models.CandidateArgument{
			ID:               "xxx",
			FirstName:        "dhany",
			LastName:         "aritonang",
			PhoneNumber:      "080821",
			TelpNumber:       "080821",
			Address:          "simanabun",
			ProfilPictureUrl: "gambar.jpg",
		},
		onCheckCandidateById: func(mock *mocks.MockCandidateRepository) {
			mock.EXPECT().CheckCandidateById(gomock.Any(), gomock.Any()).Return(
				models.Candidate{
					ID:    "xxx",
					Email: "dhany@gmail.com",
				},
				nil,
			)
		},
		onUpdateCandidate: func(mock *mocks.MockCandidateRepository) {
			mock.EXPECT().UpdateCandidate(gomock.Any(), gomock.Any()).Return(
				errors.New("internal server error"),
			)
		},
	})

	testCases = append(testCases, testCase{
		desc:      "success update candidate",
		wantError: false,
		input: models.CandidateArgument{
			ID:               "xxx",
			FirstName:        "dhany",
			LastName:         "aritonang",
			PhoneNumber:      "080821",
			TelpNumber:       "080821",
			Address:          "simanabun",
			ProfilPictureUrl: "gambar.jpg",
		},
		onCheckCandidateById: func(mock *mocks.MockCandidateRepository) {
			mock.EXPECT().CheckCandidateById(gomock.Any(), gomock.Any()).Return(
				models.Candidate{
					ID:    "xxx",
					Email: "dhany@gmail.com",
				},
				nil,
			)
		},
		onUpdateCandidate: func(mock *mocks.MockCandidateRepository) {
			mock.EXPECT().UpdateCandidate(gomock.Any(), gomock.Any()).Return(
				nil,
			)
		},
		onGetDetailCandidate: func(mock *mocks.MockCandidateRepository) {
			mock.EXPECT().GetDetailCandidate(gomock.Any(), gomock.Any()).Return(
				models.Candidate{
					ID:               "xxx",
					FirstName:        "dhany",
					LastName:         "aritonang",
					PhoneNumber:      "080821",
					TelpNumber:       "080821",
					Address:          "simanabun",
					ProfilPictureUrl: "gambar.jpg",
				},
				nil,
			)
		},
	})

	for _, tC := range testCases {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		repo := mocks.NewMockCandidateRepository(mockCtrl)
		usecase := candidateUsecase{repo}

		if tC.onCheckCandidateById != nil {
			tC.onCheckCandidateById(repo)
		}
		if tC.onUpdateCandidate != nil {
			tC.onUpdateCandidate(repo)
		}
		if tC.onGetDetailCandidate != nil {
			tC.onGetDetailCandidate(repo)
		}

		candidate, err := usecase.UpdateCandidate(context.Background(), tC.input)

		t.Run(tC.desc, func(t *testing.T) {
			if tC.wantError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, candidate.ID, tC.input.ID)
			}
		})
	}
}
