use std::env;
use std::sync::{Mutex, OnceLock};
use task_tracker::operations::add::add_task;
use task_tracker::operations::delete::delete_task;
use task_tracker::storage::load_tasks;
use tempfile::tempdir;

static TEST_MUTEX: OnceLock<Mutex<()>> = OnceLock::new();

fn get_test_lock() -> std::sync::MutexGuard<'static, ()> {
    TEST_MUTEX.get_or_init(|| Mutex::new(())).lock().unwrap()
}

#[test]
fn test_delete_task() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    add_task("Task 1", None).unwrap();
    add_task("Task 2", None).unwrap();
    add_task("Task 3", None).unwrap();

    delete_task(2).unwrap();

    let task_list = load_tasks().unwrap();
    assert_eq!(task_list.tasks.len(), 2);
    assert_eq!(task_list.tasks[0].id, 1);
    assert_eq!(task_list.tasks[1].id, 3);
}

#[test]
fn test_delete_all_tasks_resets_id() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    add_task("Task 1", None).unwrap();
    add_task("Task 2", None).unwrap();

    let task_list = load_tasks().unwrap();
    assert_eq!(task_list.tasks.len(), 2);
    assert_eq!(task_list.next_id, 3);

    delete_task(1).unwrap();

    let task_list = load_tasks().unwrap();
    assert_eq!(task_list.tasks.len(), 1);

    delete_task(2).unwrap();

    let task_list = load_tasks().unwrap();
    assert_eq!(task_list.tasks.len(), 0);
    assert_eq!(task_list.next_id, 1);
}

#[test]
fn test_delete_nonexistent_task() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    let result = delete_task(999);
    assert!(result.is_err());
}
