package telegram

type TelegramUploadResponse struct {
	Ok     bool                 `json:"ok"`
	Result TelegramUploadResult `json:"result"`
}

type TelegramDownloadResponse struct {
	Ok     bool                     `json:"ok"`
	Result TelegramResponseDocument `json:"result"`
}

type TelegramUploadResult struct {
	Document TelegramResponseDocument `json:"document"`
}

type TelegramResponseDocument struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FilePath     string `json:"file_path"`
}

type TelegramDownloadError struct {
	Ok        bool    `json:"ok"`
	ErrorCode float64 `json:"error_code"`
	Descr     string  `json:"description"`
}

type Telegram struct {
	Token  string `json:"token"`
	ChatId string `json:"chat_id"`
}
