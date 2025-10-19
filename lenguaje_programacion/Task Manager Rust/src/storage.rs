use crate::types::TaskList;
use std::fs;
use std::io::{self, ErrorKind};
use std::path::Path;

const STORAGE_FILE: &str = "tasks.json";

pub fn load_tasks() -> io::Result<TaskList> {
    if !Path::new(STORAGE_FILE).exists() {
        return Ok(TaskList::default());
    }

    let content = fs::read_to_string(STORAGE_FILE)?;

    if content.trim().is_empty() {
        return Ok(TaskList::default());
    }

    serde_json::from_str(&content).map_err(|e| io::Error::new(ErrorKind::InvalidData, e))
}

pub fn save_tasks(task_list: &TaskList) -> io::Result<()> {
    let json = serde_json::to_string_pretty(task_list)
        .map_err(|e| io::Error::new(ErrorKind::InvalidData, e))?;

    fs::write(STORAGE_FILE, json)?;
    Ok(())
}
