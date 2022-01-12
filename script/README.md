# Init Problem Scipt

## Use Go (Only Create Empty file)
只能在 repo root 位置執行（相對於指令執行位置）
```sh
# in project root path run
go run script/init_problem.go -title="160. Intersection of Two Linked Lists"

# print
create folder:  ./problems/0160.Intersection-of-Two-Linked-Lists
create file:  ./problems/0160.Intersection-of-Two-Linked-Lists/0160.go
create test file:  ./problems/0160.Intersection-of-Two-Linked-Lists/0160_test.go
```

## Use Ansible (Build File With jinja2 Template)
可以在任何位置執行（相對於 ansible.yaml 檔案位置）
```sh
ansible-playbook ./script/init_problem_ansible.yml -e "title='160. Intersection of Two Linked Lists'" -v
```