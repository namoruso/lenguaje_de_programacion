use std::env;
use std::sync::{Mutex, OnceLock};
use task_tracker::operations::add::add_task;
use task_tracker::operations::list::{
    list_all_tasks, list_done_tasks, list_in_progress_tasks, list_todo_tasks,
};
use task_tracker::operations::status::{mark_done, mark_in_progress};
use tempfile::tempdir;

static TEST_MUTEX: OnceLock<Mutex<()>> = OnceLock::new();

fn get_test_lock() -> std::sync::MutexGuard<'static, ()> {
    TEST_MUTEX.get_or_init(|| Mutex::new(())).lock().unwrap()
}

#[test]
fn test_list_all_tasks() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    add_task("Task 1", Some("Description 1".to_string())).unwrap();
    add_task("Task 2", None).unwrap();

    let result = list_all_tasks();
    assert!(result.is_ok());
}

#[test]
fn test_list_filtered_tasks() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    add_task("Todo task", None).unwrap();
    add_task("Progress task", Some("In progress".to_string())).unwrap();
    add_task("Done task", Some("Completed".to_string())).unwrap();

    mark_in_progress(2).unwrap();
    mark_done(3).unwrap();

    assert!(list_todo_tasks().is_ok());
    assert!(list_in_progress_tasks().is_ok());
    assert!(list_done_tasks().is_ok());
}
