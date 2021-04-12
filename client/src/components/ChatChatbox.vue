<template>
  <b-container>
    <b-card>
      <div class="chatbox">
        <b-row
          class="no-gutters"
          align-h="start"
          v-for="(message, inx) in chatHistory"
          :key="inx"
        >
          <b-col class="no-gutters" cols="8">
            <div>
              <p class="received-chat">{{ message }}</p>
            </div>
          </b-col>
        </b-row>
      </div>
      <chat-message-box :socket="socket"></chat-message-box>
    </b-card>
  </b-container>
</template>

<script>
import ChatMessageBox from "./ChatMessagebox.vue";
export default {
  props: ["socket"],
  components: {
    ChatMessageBox,
  },
  data() {
    return {
      chatHistory: [],
    };
  },
  mounted() {
    if (this.socket) {
      this.socket.on("/message", (message) => {
        if (message) {
          this.chatHistory.push(message.trim());
        }
      });
    }
  },
};
</script>

<style scoped>
.container {
  height: calc(100% - 100px);
}
.chatbox {
  padding: 10px;
  height: calc(100% - 35px);
  overflow-y: auto;
  background-image: url("https://i0.wp.com/www.tortugacreative.com/wp-content/uploads/social-media-background-dark-2.png?ssl=1");
}
.sender-name {
  margin: 0px;
}
.chat {
  background-color: lightgreen;
  padding: 10px;
  border-top-left-radius: 10px;
  border-bottom-left-radius: 10px;
  border-bottom-right-radius: 10px;
}
.received-chat {
  background-color: lightblue;
  padding: 10px;
  border-top-right-radius: 10px;
  border-bottom-left-radius: 10px;
  border-bottom-right-radius: 10px;
}
</style>
