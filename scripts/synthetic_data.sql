INSERT INTO users (role, group_id, name) VALUES
    (0, 101, 'John Doe'),
    (0, 101, 'Jane Doe'),
    (0, 102, 'Alice Smith'),
    (1, 102, 'Bob Johnson'),
    (1, 103, 'Eve Brown'),
    (1, 103, 'Charlie Davis');

INSERT INTO tasks (group_id, name, description) VALUES
    (101, 'Task 1', 'Description for Task 1'),
    (101, 'Task 2', 'Description for Task 2'),
    (102, 'Task 3', 'Description for Task 3'),
    (102, 'Task 4', 'Description for Task 4'),
    (103, 'Task 5', 'Description for Task 5'),
    (103, 'Task 6', 'Description for Task 6');

INSERT INTO assignments (task_id, user_id, status, feedback, file) VALUES
    (1, 1, 1, 'Good job!', 'file1.txt'),
    (2, 1, 0, 'Incomplete', 'file2.txt'),
    (3, 2, 1, 'Well done!', 'file3.txt'),
    (4, 2, 0, 'Incomplete', 'file4.txt'),
    (5, 3, 1, 'Excellent!', 'file5.txt'),
    (6, 3, 1, 'Satisfactory', 'file6.txt');