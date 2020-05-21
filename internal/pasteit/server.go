package pasteit

import (
	"context"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/twitchtv/twirp"

	pb "github.com/sarathsp06/pasteit/rpc/server"
)

var validationError error

// Paste represents a data to be pasted
type Paste struct {
	Title     string
	ID        uuid.UUID
	Body      string
	Headers   map[string]string
	createdAt time.Time
	TTL       time.Duration
}

type store interface {
	Store(context.Context, Paste) (Paste, error)
}

type server struct {
	maxContent int
	minContent int
	store      store
}

// NewServer creates an instance of
func NewServer(maxContent int, minContent int, store store) pb.Paster {
	return &server{
		maxContent: maxContent,
		minContent: minContent,
		store:      store,
	}
}

func (s *server) Paste(ctx context.Context, pr *pb.PasteRequest) (*pb.PasteResponse, error) {
	if len(pr.Content) > s.maxContent || len(pr.Content) < s.minContent {
		return nil, twirp.InvalidArgumentError("Content", "Invalid content size")
	}

	paste := Paste{
		Body:    pr.Content,
		Headers: pr.Headers,
		TTL:     time.Second * time.Duration(pr.Ttl),
	}

	paste, err := s.store.Store(ctx, paste)
	if err := castError(err); err != nil {
		return nil, err
	}
	return &pb.PasteResponse{
		Uuid:      paste.ID.String(),
		CreatedAt: paste.createdAt.Unix(),
	}, err
}

func castError(err error) error {
	if err == nil {
		return nil
	}
	switch {
	case errors.Is(err, validationError):
		return twirp.NewError(twirp.InvalidArgument, "Validation error")
	case errors.Is(err, context.Canceled):
		return twirp.NewError(twirp.Canceled, "")
	case errors.Is(err, context.DeadlineExceeded):
		return twirp.NewError(twirp.Canceled, "")
	default:
		return twirp.InternalErrorWith(err)
	}
}
