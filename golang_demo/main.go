package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	var transport string
	flag.StringVar(&transport, "t", "stdio", "Transport type (stdio or sse)")
	flag.StringVar(
		&transport,
		"transport",
		"stdio",
		"Transport type (stdio or sse)",
	)
	flag.Parse()

	// Create MCP server
	s := server.NewMCPServer(
		"Demo 🚀",
		"1.0.0",
	)
	// Add tool
	tool := mcp.NewTool("add",
		mcp.WithDescription("Add two numbers"),
		mcp.WithNumber("a",
			mcp.Required(),
			mcp.Description("First number"),
		),
		mcp.WithNumber("b",
			mcp.Required(),
			mcp.Description("Second number"),
		),
	)
	// Add tool handler
	s.AddTool(tool, addHandler)

	switch transport {
	case "stdio":
		if err := server.ServeStdio(s); err != nil {
			log.Fatalf("Server error: %v\n", err)
		}
	case "sse":
		sse := server.NewSSEServer(s, server.WithBaseURL("http://localhost:8080"))
		log.Printf("SSE server listening on :8080")
		if err := sse.Start(":8080"); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	default:
		log.Fatalf(
			"Invalid transport type: %s. Must be 'stdio' or 'sse'",
			transport,
		)
	}
}

func addHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	a, ok := request.Params.Arguments["a"].(float64)
	if !ok {
		return nil, errors.New("a must be a number")
	}
	b, ok := request.Params.Arguments["b"].(float64)
	if !ok {
		return nil, errors.New("b must be a number")
	}
	res := fmt.Sprintf("%g", a+b)
	return mcp.NewToolResultText(res), nil
}
