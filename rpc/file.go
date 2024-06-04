package rpc

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"nakama-task/repo"
	"os"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

type getFileContentsPayload struct {
	Type    *string `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
	Hash    *string `json:"hash,omitempty"`
}

type getFileContentsResp struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Hash    string `json:"hash"`
	Content string `json:"content"`
}

// GetFileContents that reads a file based on input parameters, calculates its hash, and returns the file content if the hash matches.
func GetFileContents(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	var p getFileContentsPayload

	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		return "", fmt.Errorf("unmarshal payload: %w", err)
	}

	userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		return "", errors.New("unable to retrieve userID")
	}

	if _, err := repo.SaveUserRequestDetails(ctx, db, userID, "GetFileContents", time.Now()); err != nil {
		return "", fmt.Errorf("save request details: %w", err)
	}

	if p.Type == nil {
		*p.Type = "core"
	}

	if p.Version == nil {
		*p.Version = "0.0.0"
	}

	if p.Hash == nil {
		*p.Hash = "null"
	}

	file, err := os.Open(fmt.Sprintf("data/files/%s/%s.json", *p.Type, *p.Version))
	if err != nil {
		return "", errors.New("file not found")
	}

	defer file.Close()

	hash, err := generateMD5Hash(fmt.Sprintf("data/files/%s/%s.json", *p.Type, *p.Version))
	if err != nil {
		return "", errors.New("generate md5 hash")
	}

	// if hash values do not match return null json.
	if hash != *p.Hash {
		return "null", nil
	}

	// Read the file contents
	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("Error reading file contents: %w", err)
	}

	b, err := json.Marshal(getFileContentsResp{
		Type:    *p.Type,
		Version: *p.Version,
		Hash:    *p.Hash,
		Content: string(content),
	})
	if err != nil {
		return "", fmt.Errorf("marshal response: %w", err)
	}

	return string(b), nil
}

// helpers

/*
generateMD5Hash calculates the MD5 hash of the file located at the given file path.

Parameters:
- filePath: a string representing the path to the file for which the MD5 hash needs to be calculated.

Returns:
- string: the hexadecimal representation of the MD5 hash.
- error: an error if any occurred during the process, such as file not found or failure to read the file.

Example:
hash, err := generateMD5Hash("/path/to/file.txt")

	if err != nil {
	    // handle the error
	}

fmt.Println(hash) // Output: e4d909c290d0fb1ca068ffaddf22cbd0
*/
func generateMD5Hash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", errors.New("file not found")
	}

	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
