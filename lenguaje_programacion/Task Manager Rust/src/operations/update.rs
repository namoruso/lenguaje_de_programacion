use crate::storage::{load_tasks, save_tasks};
use chrono::Utc;
use std::io::{self, Error, ErrorKind};

pub fn update_task(
    id: u32,
    new_title: Option<&str>,
    new_description: Option<String>,
) -> io::Result<()> {
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

    if let Some(title) = new_title {
        task.title = title.to_string();
    }

    task.description = new_description;
    task.updated_at = Utc::now();

    save_tasks(&task_list)?;

    println!("✓ Task {} updated successfully", id);
    Ok(())
}
