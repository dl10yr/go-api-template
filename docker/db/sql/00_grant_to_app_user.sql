CREATE USER 'todo-app'@'%' IDENTIFIED BY 'password';
GRANT SELECT,INSERT,UPDATE,DELETE ON todo.* TO 'todo-app'@'%';