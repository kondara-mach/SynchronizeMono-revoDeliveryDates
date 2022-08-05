package database

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestJobBookRepository_FetchAll(t *testing.T) {
	err_read := godotenv.Load(`../../.env`)
	if err_read != nil {
		os.Exit(1)
	}

	logger, _ := zap.NewDevelopment()

	tests := []struct {
		name string
		r    *JobBookRepository
		want struct {
			workNumRex           string
			canEmptyDeliveryDate bool
		}
	}{
		{
			name: "正常系_M作業から作業Noと納期が取得できること",
			r: NewJobBookRepository(
				logger.Sugar(),
				OrderDbPram{
					os.Getenv("DB_SERVER"),
					os.Getenv("DB_NAME"),
					os.Getenv("DB_USER"),
					os.Getenv("DB_PASS"),
				},
			),
			want: struct {
				workNumRex           string
				canEmptyDeliveryDate bool
			}{
				workNumRex:           `X?[0-9]{2}[A-Z]-[0-9]{1,4}`,
				canEmptyDeliveryDate: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.r.FetchAll()
			assert.NotEmpty(t, got)
			assert.NotEqual(t, 0, len(got))
			for _, v := range got {
				assert.True(
					t,
					regexp.MustCompile(tt.want.workNumRex).
						Match(
							[]byte(v.WorkedNumber),
						),
					fmt.Sprintf("作業No(%v)の形式の検証", v.WorkedNumber),
				)
				if tt.want.canEmptyDeliveryDate {
					assert.Fail(t, "このテストケースは想定外")
				} else {
					assert.NotEmpty(t, v.DeliveryDate, "Emptyでないこと")
				}
			}
		})
	}
}
