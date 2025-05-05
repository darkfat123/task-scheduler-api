CREATE TABLE tasks (
  id SERIAL PRIMARY KEY,
  code VARCHAR(20) UNIQUE,
  name VARCHAR(100),
  frequency_date VARCHAR(20),
  frequency_time VARCHAR(20),
  next_run_at TIMESTAMP,
  last_run_at TIMESTAMP,
  max_retries INT,
  status VARCHAR(10) NOT NULL DEFAULT 'Pending',
  is_enabled BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE tasks_runs (
  id SERIAL PRIMARY KEY,
  task_id INT REFERENCES tasks(id) ON DELETE CASCADE,
  run_at TIMESTAMP DEFAULT NOW(),           
  finished_at TIMESTAMP,                    
  status VARCHAR(10),                       
  retry_count INT DEFAULT 0,                
  error_message TEXT,                       
  duration_seconds INT                      
);
