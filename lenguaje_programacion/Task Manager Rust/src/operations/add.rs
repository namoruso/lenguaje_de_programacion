use crate::storage::{load_tasks, save_tasks};
use crate::types::Task;
use std::io;

pub fn add_task(title: &str, description: Option<String>) -> io::Result<()> {
    let mut task_list = load_tasks()?;

    let task = Task::new(task_list.next_id, title.to_string(), description);
    let task_id = task.id;

    task_list.tasks.push(task);
    task_list.next_id += 1;

    save_tasks(&task_list)?;

    println!("✓ Task added successfully (ID: {})", task_id);
    Ok(())
}
