#!/bin/bash

# IP Lookup App Build Script

set -e

echo "🚀 Building IP Lookup Application..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or later."
    exit 1
fi

# Check if Node.js is installed
if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js 20 or later."
    exit 1
fi

echo "📦 Building frontend..."
cd frontend

# Install dependencies
echo "   Installing npm dependencies..."
npm install

# Build frontend
echo "   Building Vue.js application..."
npm run build

cd ..

echo "🔨 Building Go binary..."
go build -tags embed -o ip-lookup-app

echo "✅ Build complete!"
echo ""
echo "🎉 Your IP Lookup application is ready!"
echo "   Run: ./ip-lookup-app"
echo "   Then open: http://localhost:8080"
echo ""

