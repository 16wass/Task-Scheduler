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
  }
}
</script>

<template>
  <div>
    <h1> Task manager </h1>
    <ul>
      <li v-for=" task in tasks" : key="task.id">{{task.title}}} </li>
    </ul>
    <form>
      <input v-model="new.title" placeholder="Task Title"/>
      <input v-model="new.description" placeholder="Description" />
      <button type = "submit" > Add Task </button>
    </form>
  </div>
</template>

<style scoped>
</style>
