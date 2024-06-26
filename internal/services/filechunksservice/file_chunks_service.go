package repositorysyncservice

import (
	"context"
	"os"

	"github.com/grantchen2003/insight/repository/internal/services/filechunksservice/pb"

	"google.golang.org/grpc"
)

type FileChunk struct {
	UserId         string
	FilePath       string
	Content        []byte
	ChunkIndex     int
	NumTotalChunks int
}

type FileChunkSaveStatus struct {
	FilePath         string
	IsLastSavedChunk bool
}

func CreateFileChunks(userId string, fileChunks []FileChunk) ([]FileChunkSaveStatus, error) {
	conn, err := grpc.Dial(os.Getenv("FILE_CHUNKS_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var pbFileChunks []*pb.FileChunkPayload

	for _, fileChunk := range fileChunks {
		pbFileChunks = append(pbFileChunks, &pb.FileChunkPayload{
			UserId:         userId,
			FilePath:       fileChunk.FilePath,
			Content:        fileChunk.Content,
			ChunkIndex:     int32(fileChunk.ChunkIndex),
			NumTotalChunks: int32(fileChunk.NumTotalChunks),
		})
	}

	client := pb.NewFileChunksServiceClient(conn)
	response, err := client.CreateFileChunks(context.Background(), &pb.CreateFileChunksRequest{
		FileChunkPayloads: pbFileChunks,
	})

	if err != nil {
		return nil, err
	}

	var fileChunkSaveStatuses []FileChunkSaveStatus

	for _, pbFileChunkStatus := range response.FileChunkStatuses {
		fileChunkSaveStatuses = append(fileChunkSaveStatuses, FileChunkSaveStatus{
			FilePath:         pbFileChunkStatus.FilePath,
			IsLastSavedChunk: pbFileChunkStatus.IsLastSavedChunk,
		})
	}

	return fileChunkSaveStatuses, nil

}
