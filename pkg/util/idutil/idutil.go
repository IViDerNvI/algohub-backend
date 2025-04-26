package idutil

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

func SnowflakeID() uint {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	return uint(node.Generate().Int64())
}

func UUID() string {
	return uuid.New().String()
}
