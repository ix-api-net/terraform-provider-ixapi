provider "ixapi" {
    api = "http://localhost:8000/api/v2"
    api_key = "abcdef1234567890"

    # If empty, the api_secret will be read from the environment
    # variable: $IX_API_SECRET
    # api_secret = "..."
}
