package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
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

	/*
		Start the stdio server
	*/
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

	/*
		Start the sse server
	*/
	// sse := server.NewSSEServer(s, server.WithBaseURL("http://localhost:8080"))
	// if err := sse.Start(":8080"); err != nil {
	// 	fmt.Printf("Server error: %v\n", err)
	// }
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
