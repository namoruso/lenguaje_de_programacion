use crate::storage::{load_tasks, save_tasks};
use crate::types::TaskStatus;
use chrono::Utc;
use std::io::{self, Error, ErrorKind};

pub fn mark_in_progress(id: u32) -> io::Result<()> {
    change_status(id, TaskStatus::InProgress)
}

pub fn mark_done(id: u32) -> io::Result<()> {
    change_status(id, TaskStatus::Done)
}

fn change_status(id: u32, new_status: TaskStatus) -> io::Result<()> {
    let mut task_list = load_tasks()?;

    let task = task_list
        .tasks
        .iter_mut()
        .find(|t| t.id == id)
        .ok_or_else(|| {
            Error::new(
                ErrorKind::NotFound,
                format!("Task with ID {} not found", id),
            )
        })?;

    task.status = new_status.clone();
    task.updated_at = Utc::now();

    save_tasks(&task_list)?;

    println!("✓ Task {} marked as {}", id, new_status.as_str());
    Ok(())
}
