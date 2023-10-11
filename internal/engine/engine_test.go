package engine

import (
	"fmt"
	"slotengine/internal/mocks"
	"slotengine/internal/reel"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlay(t *testing.T) {
	tests := []struct {
		name                 string
		bet                  float64
		expectedMatrix       [3][4]string
		expectedWin          float64
		expectedSpinResult   []string
		expectedError        error
		spinError            error
		calculatePayoutError error
	}{
		{
			name: "PredictableOutcome",
			bet:  5.0,
			expectedMatrix: [3][4]string{
				{"A", "A", "A", "A"},
				{"K", "K", "K", "K"},
				{"Q", "Q", "Q", "Q"},
			},
			expectedWin:        100.0,
			expectedSpinResult: []string{"A", "K", "Q"},
			expectedError:      nil,
		},
		{
			name:               "Expects a spin error",
			bet:                5.0,
			expectedMatrix:     [3][4]string{},
			expectedWin:        0,
			expectedSpinResult: []string{"A", "K", "Q"},
			expectedError:      fmt.Errorf("error"),
			spinError:          fmt.Errorf("error"),
		},
		{
			name: "Expects a calculatePayoutError error",
			bet:  5.0,
			expectedMatrix: [3][4]string{
				{"A", "A", "A", "A"},
				{"K", "K", "K", "K"},
				{"Q", "Q", "Q", "Q"},
			},
			expectedWin:          0,
			expectedSpinResult:   []string{"A", "K", "Q"},
			expectedError:        fmt.Errorf("error"),
			calculatePayoutError: fmt.Errorf("error"),
		},
		{
			name:               "Bet less than 0",
			bet:                -1,
			expectedMatrix:     [3][4]string{},
			expectedWin:        0,
			expectedSpinResult: []string{"A", "K", "Q"},
			expectedError:      fmt.Errorf("bet must be greater than 0"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Mock Reel
			mockReel := &mocks.ReelMock{}
			mockReel.SpinFunc = func() ([]string, error) {
				return tt.expectedSpinResult, tt.spinError
			}

			mockReels := make([]reel.Reel, 4)
			for i := range mockReels {
				mockReels[i] = mockReel
			}

			// Mock PayTable
			mockPayTable := &mocks.PayTableMock{}
			mockPayTable.CalculatePayoutFunc = func(matrix [3][4]string, bet float64) (float64, error) {
				if matrix[0][0] == "A" && matrix[0][1] == "A" && matrix[0][2] == "A" && matrix[0][3] == "A" {
					return 100.0, tt.calculatePayoutError
				}
				return 0, nil
			}

			engine := NewBasicGameEngine(mockReels, mockPayTable)

			matrix, win, err := engine.Play(tt.bet)

			assert.Equal(t, tt.expectedError, err, "The error does not match the expected error")
			assert.Equal(t, tt.expectedMatrix, matrix, "The matrix does not match expected values")
			assert.Equal(t, tt.expectedWin, win, "The win amount does not match the expected value")
		})
	}
}
