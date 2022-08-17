package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCheckSessions(t *testing.T) {
	type args struct {
		Sessions []Session
	}
	tests := []struct {
		name string
		args args
		want []sessionConfict
	}{
		{
			name: "prtbiiisdf hrty",
			args: args{
				Sessions: []Session{
					{
						name:       "superman",
						hall_id:    1,
						startTime:  time.Date(2022, time.August, 15, 17, 35, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 18, 35, 0, 0, time.UTC),
						session_id: "123",
					},
					{
						name:       "thor",
						hall_id:    2,
						startTime:  time.Date(2022, time.August, 15, 16, 30, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 17, 30, 0, 0, time.UTC),
						session_id: "sss1",
					},
					{
						name:       "joker",
						hall_id:    3,
						startTime:  time.Date(2022, time.August, 15, 19, 40, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 20, 30, 0, 0, time.UTC),
						session_id: "1s3",
					},
					{
						name:       "batman",
						hall_id:    1,
						startTime:  time.Date(2022, time.August, 15, 17, 40, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 18, 20, 0, 0, time.UTC),
						session_id: "1ss3",
					},
				},
			},
			want: []sessionConfict{
				{
					ids: []string{"1ss3", "123"},
				},
			},
		},
		{
			name: "vtotito",
			args: args{
				Sessions: []Session{
					{
						name:       "superman",
						hall_id:    1,
						startTime:  time.Date(2022, time.August, 15, 17, 35, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 18, 35, 0, 0, time.UTC),
						session_id: "123",
					},
					{
						name:       "thor",
						hall_id:    2,
						startTime:  time.Date(2022, time.August, 15, 16, 30, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 17, 30, 0, 0, time.UTC),
						session_id: "sss1",
					},
					{
						name:       "joker",
						hall_id:    3,
						startTime:  time.Date(2022, time.August, 15, 19, 40, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 20, 30, 0, 0, time.UTC),
						session_id: "133",
					},
					{
						name:       "batman",
						hall_id:    1,
						startTime:  time.Date(2022, time.August, 15, 19, 20, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 20, 30, 0, 0, time.UTC), // if End time is less than start time
						session_id: "1s3",
					},
				},
			},
			want: nil,
		},
		{
			name: "tretiy",
			args: args{
				Sessions: []Session{
					{
						name:       "superman",
						hall_id:    1,
						startTime:  time.Date(2022, time.August, 15, 17, 35, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 18, 35, 0, 0, time.UTC),
						session_id: "123",
					},
					{
						name:       "thor",
						hall_id:    2,
						startTime:  time.Date(2022, time.August, 15, 16, 30, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 17, 30, 0, 0, time.UTC),
						session_id: "sss1",
					},
					{
						name:       "joker",
						hall_id:    3,
						startTime:  time.Date(2022, time.August, 15, 19, 40, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 20, 30, 0, 0, time.UTC),
						session_id: "133",
					},
					{
						name:       "batman",
						hall_id:    1,
						startTime:  time.Date(2022, time.August, 15, 19, 20, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 20, 30, 0, 0, time.UTC), // if End time is less than start time
						session_id: "1s3",
					},
					{
						name:       "wonderwoman",
						hall_id:    2,
						startTime:  time.Date(2022, time.August, 15, 17, 35, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 18, 35, 0, 0, time.UTC),
						session_id: "12113",
					},
					{
						name:       "terminator",
						hall_id:    2,
						startTime:  time.Date(2022, time.August, 15, 16, 30, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 17, 30, 0, 0, time.UTC),
						session_id: "ssaa1",
					},
					{
						name:       "spiderman",
						hall_id:    3,
						startTime:  time.Date(2022, time.August, 15, 19, 40, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 20, 30, 0, 0, time.UTC),
						session_id: "13sss3",
					},
					{
						name:       "mysterio",
						hall_id:    3,
						startTime:  time.Date(2022, time.August, 15, 19, 20, 0, 0, time.UTC),
						endTime:    time.Date(2022, time.August, 15, 20, 30, 0, 0, time.UTC), // HW: if End time is less than start time
						session_id: "1s3aa",
					},
				},
			},
			want: []sessionConfict{
				{
					ids: []string{"ssaa1", "sss1"},
				},
				{
					ids: []string{"133", "1s3aa"},
				},
				{
					ids: []string{"13sss3", "133"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSessions(tt.args.Sessions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckSessions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckSession(t *testing.T) {
	type args struct {
		sessions []Session
	}
	tests := []struct {
		name string
		args args
		want []sessionError
	}{
		{
			name: "gat",
			args: args{sessions: []Session{
				{
					name:       "superman",
					hall_id:    1,
					startTime:  time.Date(2022, time.August, 15, 17, 35, 0, 0, time.UTC),
					endTime:    time.Date(2022, time.August, 15, 16, 35, 0, 0, time.UTC),
					session_id: "123",
					status:     "",
				},
			},
			},
			want: []sessionError{
				{
					ids: []string{"123"},
				},
			},
		},
		{
			name: "gat",
			args: args{sessions: []Session{
				{
					name:       "superman",
					hall_id:    1,
					startTime:  time.Date(2022, time.August, 15, 17, 35, 0, 0, time.UTC),
					endTime:    time.Date(2022, time.August, 15, 18, 35, 0, 0, time.UTC),
					session_id: "123",
					status:     "",
				},
				{
					name:       "batman beyond",
					hall_id:    1,
					startTime:  time.Date(2022, time.August, 15, 18, 40, 0, 0, time.UTC),
					endTime:    time.Date(2022, time.August, 15, 19, 35, 0, 0, time.UTC),
					session_id: "123",
					status:     "",
				},
			},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSession(tt.args.sessions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
