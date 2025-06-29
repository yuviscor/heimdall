package notifier

func getStatusText(statusCode int) string {
	switch {
	case statusCode >= 200 && statusCode < 300:
		return "Success"
	case statusCode >= 300 && statusCode < 400:
		return "Redirect"
	case statusCode >= 400 && statusCode < 500:
		return "Client Error"
	case statusCode >= 500:
		return "Server Error"
	default:
		return "Unknown"
	}
}
