mod operations;
mod storage;
mod types;

use operations::{add, delete, list, status, update};
use std::env;
use std::process;

fn print_usage() {
    println!("Task Tracker CLI");
    println!("\nUsage:");
    println!("  task-tracker add <title> [description]           - Add a new task");
    println!("  task-tracker update <id> [--title <title>] [--desc <description>] - Update a task");
    println!("  task-tracker delete <id>                         - Delete a task");
    println!("  task-tracker mark-in-progress <id>               - Mark task as in progress");
    println!("  task-tracker mark-done <id>                      - Mark task as done");
    println!("  task-tracker list                                - List all tasks");
    println!("  task-tracker list done                           - List done tasks");
    println!("  task-tracker list todo                           - List todo tasks");
    println!("  task-tracker list in-progress                    - List in-progress tasks");
    println!("\nExamples:");
    println!("  task-tracker add \"Buy groceries\"");
    println!("  task-tracker add \"Buy groceries\" \"Milk, eggs, and bread\"");
    println!("  task-tracker update 1 --title \"Buy groceries and cook\"");
    println!("  task-tracker update 1 --desc \"Milk, eggs, bread, and vegetables\"");
    println!("  task-tracker update 1 --title \"New title\" --desc \"New description\"");
    println!("  task-tracker mark-in-progress 1");
    println!("  task-tracker mark-done 1");
    println!("  task-tracker delete 1");
}

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        print_usage();
        process::exit(1);
    }

    let result = match args[1].as_str() {
        "add" => {
            if args.len() < 3 {
                eprintln!("Error: Missing task title");
                print_usage();
                process::exit(1);
            }
            let title = &args[2];
            let description = if args.len() > 3 {
                Some(args[3..].join(" "))
            } else {
                None
            };
            add::add_task(title, description)
        }
        "update" => {
            if args.len() < 3 {
                eprintln!("Error: Missing task ID");
                print_usage();
                process::exit(1);
            }
            let id: u32 = match args[2].parse() {
                Ok(id) => id,
                Err(_) => {
                    eprintln!("Error: Invalid task ID");
                    process::exit(1);
                }
            };

            let mut new_title: Option<&str> = None;
            let mut new_description: Option<String> = None;
            let mut i = 3;

            while i < args.len() {
                match args[i].as_str() {
                    "--title" => {
                        if i + 1 < args.len() {
                            new_title = Some(&args[i + 1]);
                            i += 2;
                        } else {
                            eprintln!("Error: --title requires a value");
                            process::exit(1);
                        }
                    }
                    "--desc" => {
                        if i + 1 < args.len() {
                            new_description = Some(args[i + 1].clone());
                            i += 2;
                        } else {
                            eprintln!("Error: --desc requires a value");
                            process::exit(1);
                        }
                    }
                    _ => {
                        eprintln!("Error: Unknown option '{}'", args[i]);
                        print_usage();
                        process::exit(1);
                    }
                }
            }

            if new_title.is_none() && new_description.is_none() {
                eprintln!("Error: Must specify at least --title or --desc");
                print_usage();
                process::exit(1);
            }

            update::update_task(id, new_title, new_description)
        }
        "delete" => {
            if args.len() < 3 {
                eprintln!("Error: Missing task ID");
                print_usage();
                process::exit(1);
            }
            let id: u32 = match args[2].parse() {
                Ok(id) => id,
                Err(_) => {
                    eprintln!("Error: Invalid task ID");
                    process::exit(1);
                }
            };
            delete::delete_task(id)
        }
        "mark-in-progress" => {
            if args.len() < 3 {
                eprintln!("Error: Missing task ID");
                print_usage();
                process::exit(1);
            }
            let id: u32 = match args[2].parse() {
                Ok(id) => id,
                Err(_) => {
                    eprintln!("Error: Invalid task ID");
                    process::exit(1);
                }
            };
            status::mark_in_progress(id)
        }
        "mark-done" => {
            if args.len() < 3 {
                eprintln!("Error: Missing task ID");
                print_usage();
                process::exit(1);
            }
            let id: u32 = match args[2].parse() {
                Ok(id) => id,
                Err(_) => {
                    eprintln!("Error: Invalid task ID");
                    process::exit(1);
                }
            };
            status::mark_done(id)
        }
        "list" => {
            if args.len() == 2 {
                list::list_all_tasks()
            } else {
                match args[2].as_str() {
                    "done" => list::list_done_tasks(),
                    "todo" => list::list_todo_tasks(),
                    "in-progress" => list::list_in_progress_tasks(),
                    _ => {
                        eprintln!("Error: Invalid list filter. Use: done, todo, or in-progress");
                        print_usage();
                        process::exit(1);
                    }
                }
            }
        }
        "help" | "--help" | "-h" => {
            print_usage();
            process::exit(0);
        }
        _ => {
            eprintln!("Error: Unknown command '{}'", args[1]);
            print_usage();
            process::exit(1);
        }
    };

    if let Err(e) = result {
        eprintln!("Error: {}", e);
        process::exit(1);
    }
}
