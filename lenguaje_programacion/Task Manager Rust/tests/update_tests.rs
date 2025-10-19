use std::env;
use std::sync::{Mutex, OnceLock};
use task_tracker::operations::add::add_task;
use task_tracker::operations::update::update_task;
use task_tracker::storage::load_tasks;
use tempfile::tempdir;

static TEST_MUTEX: OnceLock<Mutex<()>> = OnceLock::new();

fn get_test_lock() -> std::sync::MutexGuard<'static, ()> {
    TEST_MUTEX.get_or_init(|| Mutex::new(())).lock().unwrap()
}

#[test]
fn test_update_task_title() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    add_task("Original title", Some("Original description".to_string())).unwrap();
    update_task(
        1,
        Some("Updated title"),
        Some("Original description".to_string()),
    )
    .unwrap();

    let task_list = load_tasks().unwrap();
    assert_eq!(task_list.tasks[0].title, "Updated title");
}

#[test]
fn test_update_task_description() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    add_task("Title", None).unwrap();
    update_task(1, None, Some("New description".to_string())).unwrap();

    let task_list = load_tasks().unwrap();
    assert_eq!(
        task_list.tasks[0].description,
        Some("New description".to_string())
    );
}

#[test]
fn test_update_nonexistent_task() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    let result = update_task(999, Some("Title"), None);
    assert!(result.is_err());
}
