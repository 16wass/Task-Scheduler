<script >
export default {
  data(){
    return{
      tasks:[],
      newTask:{
        title: "",
        description :""
      },
      editingTask:{}
    };
  },
  methods: {
    async fetchTasks() {
      const res = await fetch('http://localhost:8080/tasks');
      this.tasks = await res.json();
    },
    async addTask() {
      await fetch('http://localhost:8080/tasks/create', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(this.newTask)
      });
      this.newTask = { title: '', description: '' };
      this.fetchTasks();
    },
    async deleteTask(Id) {
      await fetch(`http://localhost:8080/tasks/${Id}`, {
        method: 'DELETE',
      })
        .then((response) => {
          if (!response.ok) {
            throw new Error("Failed to delete task");
          }
          this.tasks = this.tasks.filter((task) => task.id !== id);
          this.successMessage = "Task deleted successfully!";
          this.errorMessage = null;
        })
        .catch((error) => {
          this.errorMessage = "Error deleting task: " + error.message;
        });

      this.fetchTasks();
    },
    async startEdit(task) {
      this.editingTask = { ...task };
    },
    async updateTask(){
      if (this.editingTask) {
        const {id, title, description} = this.editingTask;
        try {
          const response = await fetch(`http://localhost:8080/tasks/${id}`, {
            method: 'PUT',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({title, description})
          });
          if (!response.ok) {
            throw new Error("Failed to update task");
          }
          this.fetchTasks();
          this.editingTask = null;
          this.successMessage = "Task updated successfully!";
          this.errorMessage = null;
        } catch (error) {
          this.errorMessage = "Error updating task: " + error.message;
        }
      } else {
        this.errorMessage = "No task selected for editing!";
      }
    },
    async cancelEdit(){
      this.editingTask ={}
    }

  },
  mounted() {
    this.fetchTasks();
  }
}
</script>

<template>
  <div>
    <h1>Task Manager</h1>
    <ul>
      <li v-for="task in tasks" :key="task.id">{{ task.title }}</li>
    </ul>
    <!-- Adding Task Form -->
    <form @submit.prevent="addTask">
      <input v-model="newTask.title" placeholder="Task Title" />
      <input v-model="newTask.description" placeholder="Description" />
      <button type="submit">Add Task</button>
    </form>
    <h2>Tasks</h2>
    <ul>
      <li v-for="task in tasks" :key="task.id">
        <div>
          <span><strong>{{ task.name }}</strong> - {{ task.description }}</span>
          <button @click="startEdit(task)">Edit</button>
          <button @click="deleteTask(task.id)">Delete</button>
        </div>

        <!-- Edit Task Form -->
        <div v-if="editingTask.id === task.id">
          <h3>Edit Task</h3>
          <form @submit.prevent="updateTask">
            <div>
              <label for="edit-name">Task Name:</label>
              <input
                v-model="editingTask.name"
                type="text"
                id="edit-name"
                required
              />
            </div>
            <div>
              <label for="Edit-description">Description:</label>
              <input
                v-model="editingTask.description"
                type="text"
                id="edit-description"
                required
              />
            </div>
            <button type="submit">Save</button>
            <button type="button" @click="cancelEdit">Cancel</button>
          </form>
        </div>
      </li>
    </ul>
  </div>
</template>

<style scoped>
.task-manager {
  font-family: Arial, sans-serif;
  margin: 20px;
}
form {
  margin-bottom: 20px;
}
form div {
  margin-bottom: 10px;
}
button {
  margin-left: 5px;
}
.error {
  color: red;
}
.success {
  color: green;
}
</style>



