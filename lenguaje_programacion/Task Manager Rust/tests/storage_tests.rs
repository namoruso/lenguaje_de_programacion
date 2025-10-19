use std::env;
use std::sync::{Mutex, OnceLock};
use task_tracker::storage::{load_tasks, save_tasks};
use task_tracker::types::{Task, TaskList};
use tempfile::tempdir;

static TEST_MUTEX: OnceLock<Mutex<()>> = OnceLock::new();

fn get_test_lock() -> std::sync::MutexGuard<'static, ()> {
    TEST_MUTEX.get_or_init(|| Mutex::new(())).lock().unwrap()
}

#[test]
fn test_load_tasks_creates_default_when_file_missing() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    let task_list = load_tasks().unwrap();
    assert_eq!(task_list.tasks.len(), 0);
    assert_eq!(task_list.next_id, 1);
}

#[test]
fn test_save_and_load_tasks() {
    let _lock = get_test_lock();
    let dir = tempdir().unwrap();
    env::set_current_dir(&dir).unwrap();

    let mut task_list = TaskList::default();
    task_list.tasks.push(Task::new(
        1,
        "Test task".to_string(),
        Some("Description".to_string()),
    ));
    task_list.next_id = 2;

    save_tasks(&task_list).unwrap();
    let loaded = load_tasks().unwrap();

    assert_eq!(loaded.tasks.len(), 1);
    assert_eq!(loaded.tasks[0].title, "Test task");
    assert_eq!(loaded.next_id, 2);
}
