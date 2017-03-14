package cluster_test

import (
	"testing"

	. "github.com/10gen/mongo-go-driver/cluster"
	"github.com/10gen/mongo-go-driver/model"
	"github.com/stretchr/testify/require"
)

func TestWriteSelector_ReplicaSetWithPrimary(t *testing.T) {
	t.Parallel()

	require := require.New(t)

	c := &model.Cluster{
		Kind: model.ReplicaSetWithPrimary,
		Servers: []*model.Server{
			&model.Server{
				Addr: model.Addr("localhost:27017"),
				Kind: model.RSPrimary,
			},
			&model.Server{
				Addr: model.Addr("localhost:27018"),
				Kind: model.RSSecondary,
			},
			&model.Server{
				Addr: model.Addr("localhost:27018"),
				Kind: model.RSSecondary,
			},
		},
	}

	result, err := WriteSelector()(c, c.Servers)

	require.NoError(err)
	require.Len(result, 1)
	require.Equal([]*model.Server{c.Servers[0]}, result)
}

func TestWriteSelector_ReplicaSetNoPrimary(t *testing.T) {
	t.Parallel()

	require := require.New(t)

	c := &model.Cluster{
		Kind: model.ReplicaSetNoPrimary,
		Servers: []*model.Server{
			&model.Server{
				Addr: model.Addr("localhost:27018"),
				Kind: model.RSSecondary,
			},
			&model.Server{
				Addr: model.Addr("localhost:27018"),
				Kind: model.RSSecondary,
			},
		},
	}

	result, err := WriteSelector()(c, c.Servers)

	require.NoError(err)
	require.Len(result, 0)
	require.Empty(result)
}

func TestWriteSelector_Sharded(t *testing.T) {
	t.Parallel()

	require := require.New(t)

	c := &model.Cluster{
		Kind: model.Sharded,
		Servers: []*model.Server{
			&model.Server{
				Addr: model.Addr("localhost:27018"),
				Kind: model.Mongos,
			},
			&model.Server{
				Addr: model.Addr("localhost:27018"),
				Kind: model.Mongos,
			},
		},
	}

	result, err := WriteSelector()(c, c.Servers)

	require.NoError(err)
	require.Len(result, 2)
	require.Equal([]*model.Server{c.Servers[0], c.Servers[1]}, result)
}

func TestWriteSelector_Single(t *testing.T) {
	t.Parallel()

	require := require.New(t)

	c := &model.Cluster{
		Kind: model.Single,
		Servers: []*model.Server{
			&model.Server{
				Addr: model.Addr("localhost:27018"),
				Kind: model.Standalone,
			},
		},
	}

	result, err := WriteSelector()(c, c.Servers)

	require.NoError(err)
	require.Len(result, 1)
	require.Equal([]*model.Server{c.Servers[0]}, result)
}
