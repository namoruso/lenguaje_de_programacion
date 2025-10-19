use std::env;
use std::sync::{Mutex, OnceLock};
use task_tracker::operations::add::add_task;
use task_tracker::storage::load_tasks;
use tempfile::tempdir;

static TEST_MUTEX: OnceLock<Mutex<()>> = OnceLock::new();

fn get_test_lock() -> std::sync::MutexGuard<'static, ()> {
    TEST_MUTEX.get_or_init(|| Mutex::new(())).lock().unwrap()
}

#[test]
fn test_add_task_with_description() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    add_task("Test task", Some("This is a description".to_string())).unwrap();

    let task_list = load_tasks().unwrap();
    assert_eq!(task_list.tasks.len(), 1);
    assert_eq!(task_list.tasks[0].title, "Test task");
    assert_eq!(
        task_list.tasks[0].description,
        Some("This is a description".to_string())
    );
}

#[test]
fn test_add_task_without_description() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    add_task("Test task", None).unwrap();

    let task_list = load_tasks().unwrap();
    assert_eq!(task_list.tasks.len(), 1);
    assert_eq!(task_list.tasks[0].title, "Test task");
    assert_eq!(task_list.tasks[0].description, None);
}

#[test]
fn test_add_task_increments_id() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    add_task("First", None).unwrap();
    let task_list = load_tasks().unwrap();
    assert_eq!(task_list.tasks[0].id, 1);

    add_task("Second", Some("Description".to_string())).unwrap();
    let task_list = load_tasks().unwrap();
    assert_eq!(task_list.tasks[1].id, 2);
    assert_eq!(task_list.tasks.len(), 2);
}
