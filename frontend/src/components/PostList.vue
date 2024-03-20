<template>
  <button @click="getPosts">Get posts</button>
  <ul>
    <li v-for="post in posts" :key="post.id" data-test="post">
      {{ post.title }}
    </li>
  </ul>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      posts: null,
    };
  },
  methods: {
    async getPosts() {
      console.log("DUPADUPAS ", this.test1());
      this.posts = this.test1();
    },
    test1() {
      const mockPostList = [
        { id: 1, title: "title1" },
        { id: 2, title: "title2" },
      ];
      return mockPostList;
    },
  },
};
</script>
