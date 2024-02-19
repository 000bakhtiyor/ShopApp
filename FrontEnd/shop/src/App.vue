<template>
  <div>
    <form @submit.prevent="submitForm">
      <label for="username">Username:</label>
      <input type="text" id="username" v-model="formData.username" required>
      <label for="password">Password:</label>
      <input type="password" id="password" v-model="formData.password" required>
      <label for="repeatPassword">Repeat password:</label>
      <input type="password" id="repeatPassword" v-model="formData.repeatPassword" required>
      <button type="submit">Register</button>
    </form>

    <div v-if="jsonResponse">
      <h3>Json response:</h3>
      <pre>
        {{ JSON.stringify(jsonResponse, null, 2) }}
      </pre>
    </div>
  </div>
</template>
<script>
  import axios from 'axios';

  export default {
    data() {
      return {
        formData:{
          username:'',
          password:''
        },
        jsonResponse: null,
      };
    },
    methods: {
      async submitForm() {
        try {
          const response = await axios.post('http://localhost:8080/register', this.formData);
          this.jsonResponse = response.data
        } catch (error){
          console.error('error catching:', error);
        }
        },
      },
    };
</script>