<template>
  <div class="hello">
    <h1>Welcome to Your Vue.js App <span style="color: red">Ver. 2</span></h1>
    <p>This data was fetched from: <span style="color: red">{{ dbVersion }}</span></p>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "HelloWorld",
  data() {
    return {
      dbVersion: ""
    };
  },
  methods: {
    fetchDBVersion() {
      axios.get("/api/db-version").then(v => {
        this.dbVersion = v.data.version;
      });
    }
  },
  created() {
    this.fetchDBVersion();
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
