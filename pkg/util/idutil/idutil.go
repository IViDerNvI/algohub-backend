package idutil

import "github.com/bwmarrin/snowflake"

func SnowflakeID() uint {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	return uint(node.Generate().Int64())
}
