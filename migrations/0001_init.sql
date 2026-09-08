CREATE TABLE IF NOT EXISTS logs {
    id INTEGER PRIMARY KEY AUTOINCREMENT, 
    appointment_id TEXT NOT NULL,
    client_phone TEXT, 
    trigger_type TEXT NOT NULL, -- 'booking', 'reminder', 'feedback'
    sent_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    status TEXT NOT NULL, -- 'success', 'failed'
    error_message TEXT
}

CREATE UNIQUE INDEX IF NOT EXISTS idx_appointment_id ON logs (appointment_id, trigger_type)