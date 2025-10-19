use std::env;
use std::sync::{Mutex, OnceLock};
use task_tracker::operations::add::add_task;
use task_tracker::operations::status::{mark_done, mark_in_progress};
use task_tracker::storage::load_tasks;
use task_tracker::types::TaskStatus;
use tempfile::tempdir;

static TEST_MUTEX: OnceLock<Mutex<()>> = OnceLock::new();

fn get_test_lock() -> std::sync::MutexGuard<'static, ()> {
    TEST_MUTEX.get_or_init(|| Mutex::new(())).lock().unwrap()
}

#[test]
fn test_mark_in_progress() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    add_task("Test task", None).unwrap();
    mark_in_progress(1).unwrap();

    let task_list = load_tasks().unwrap();
    assert_eq!(task_list.tasks[0].status, TaskStatus::InProgress);
}

#[test]
fn test_mark_done() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    add_task("Test task", Some("Description".to_string())).unwrap();
    mark_done(1).unwrap();

    let task_list = load_tasks().unwrap();
    assert_eq!(task_list.tasks[0].status, TaskStatus::Done);
}

#[test]
fn test_change_status_nonexistent_task() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    let result = mark_done(999);
    assert!(result.is_err());
}
