<script >
export default {
  data(){
    return{
      tasks:[],
      new:{
        title: "",
        description :""
      }
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
    <form @submit.prevent="addTask">
      <input v-model="newTask.title" placeholder="Task Title" />
      <input v-model="newTask.description" placeholder="Description" />
      <button type="submit">Add Task</button>
    </form>
  </div>
</template>

<style scoped>
</style>
