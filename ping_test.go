package wfi

import (
	"testing"
	"time"
)

func TestPingPostgres(t *testing.T) {
	if err := Up("./test", "docker-compose.database.yaml", "--build"); err != nil {
		t.Error(err)
	}
	defer Down("./test", "docker-compose.database.yaml")

	type args struct {
		host     string
		duration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "wait for database",
			args: args{
				host:     "127.0.0.1:5432",
				duration: time.Second * 15,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Ping(tt.args.host, tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPingNsq(t *testing.T) {
	if err := Up("./test", "docker-compose.nsq.yaml"); err != nil {
		t.Error(err)
	}
	defer Down("./test", "docker-compose.nsq.yaml")

	type args struct {
		host     string
		duration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "wait for nsqlookupd",
			args: args{
				host:     "localhost:4161",
				duration: time.Second * 15,
			},
			wantErr: false,
		},
		{
			name: "wait for nsqld",
			args: args{
				host:     "localhost:4151",
				duration: time.Second * 15,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Ping(tt.args.host, tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFind(t *testing.T) {
	r, err := UpWithLogs("./test", "docker-compose.kafka.yaml")
	if err != nil {
		t.Error(err)
	}
	defer Down("./test", "docker-compose.kafka.yaml")

	if err = Find("started (kafka.server.KafkaServer)", r, time.Second*10); err != nil {
		t.Errorf(`"started (kafka.server.KafkaServer)" phrase should show up within 10s, %v`, err)
	}
	if err = Find("a random message", r, time.Second*10); err == nil {
		t.Error(`'a random message' phrase should not exist`)
	}
}
