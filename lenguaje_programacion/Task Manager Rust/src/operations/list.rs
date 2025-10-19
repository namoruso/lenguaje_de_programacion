use crate::storage::load_tasks;
use crate::types::{Task, TaskStatus};
use prettytable::{format, Cell, Row, Table};
use std::io;

pub fn list_all_tasks() -> io::Result<()> {
    let task_list = load_tasks()?;
    display_tasks(&task_list.tasks, "All Tasks");
    Ok(())
}

pub fn list_done_tasks() -> io::Result<()> {
    let task_list = load_tasks()?;
    let done_tasks: Vec<&Task> = task_list
        .tasks
        .iter()
        .filter(|t| t.status == TaskStatus::Done)
        .collect();
    display_tasks(
        &done_tasks.into_iter().cloned().collect::<Vec<_>>(),
        "Done Tasks",
    );
    Ok(())
}

pub fn list_todo_tasks() -> io::Result<()> {
    let task_list = load_tasks()?;
    let todo_tasks: Vec<&Task> = task_list
        .tasks
        .iter()
        .filter(|t| t.status == TaskStatus::Todo)
        .collect();
    display_tasks(
        &todo_tasks.into_iter().cloned().collect::<Vec<_>>(),
        "Todo Tasks",
    );
    Ok(())
}

pub fn list_in_progress_tasks() -> io::Result<()> {
    let task_list = load_tasks()?;
    let in_progress_tasks: Vec<&Task> = task_list
        .tasks
        .iter()
        .filter(|t| t.status == TaskStatus::InProgress)
        .collect();
    display_tasks(
        &in_progress_tasks.into_iter().cloned().collect::<Vec<_>>(),
        "In Progress Tasks",
    );
    Ok(())
}

fn display_tasks(tasks: &[Task], title: &str) {
    if tasks.is_empty() {
        println!("\n{}: No tasks found", title);
        return;
    }

    let mut table = Table::new();
    table.set_format(*format::consts::FORMAT_BOX_CHARS);

    table.set_titles(Row::new(vec![
        Cell::new("ID"),
        Cell::new("Title"),
        Cell::new("Description"),
        Cell::new("Status"),
        Cell::new("Created At"),
        Cell::new("Updated At"),
    ]));

    for task in tasks {
        let status_cell = match task.status {
            TaskStatus::Todo => Cell::new("📋 Todo"),
            TaskStatus::InProgress => Cell::new("🔄 In Progress"),
            TaskStatus::Done => Cell::new("✅ Done"),
        };

        let description_text = match &task.description {
            Some(desc) => {
                if desc.len() > 30 {
                    format!("{}...", &desc[..27])
                } else {
                    desc.clone()
                }
            }
            None => "-".to_string(),
        };

        table.add_row(Row::new(vec![
            Cell::new(&task.id.to_string()),
            Cell::new(&task.title),
            Cell::new(&description_text),
            status_cell,
            Cell::new(&task.created_at.format("%Y-%m-%d %H:%M").to_string()),
            Cell::new(&task.updated_at.format("%Y-%m-%d %H:%M").to_string()),
        ]));
    }

    println!("\n{}", title);
    table.printstd();
}
