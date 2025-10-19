use crate::storage::{load_tasks, save_tasks};
use std::io::{self, Error, ErrorKind};

pub fn delete_task(id: u32) -> io::Result<()> {
    let mut task_list = load_tasks()?;

    let initial_len = task_list.tasks.len();
    task_list.tasks.retain(|t| t.id != id);

    if task_list.tasks.len() == initial_len {
        return Err(Error::new(
            ErrorKind::NotFound,
            format!("Task with ID {} not found", id),
        ));
    }

    if task_list.tasks.is_empty() {
        task_list.next_id = 1;
        println!("✓ Task {} deleted successfully (ID counter reset)", id);
    } else {
        println!("✓ Task {} deleted successfully", id);
    }

    save_tasks(&task_list)?;

    Ok(())
}
