import axios from "axios";
import { reactive } from "vue";

class PlayerState {
  connected: boolean = false;

  running!: boolean;
  current_idx!: number;
  queue!: number[];
  step_duration!: number;
  on_end_stop!: boolean;
  step!: number;
  updated!: number;

  private ws!: WebSocket;
  private subscribers: Map<any, () => void> = new Map<any, () => void>();

  connect() {
    this.ws = new WebSocket("/v1/ws");

    this.ws.onopen = () => {
      console.log("Connected to server");
      this.connected = true;
    };

    this.ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      this.running = data.running;
      this.current_idx = data.current_idx;
      this.queue = data.queue;
      this.step_duration = data.step_duration;
      this.on_end_stop = data.on_end_stop;
      this.step = data.step;
      this.updated = data.updated;

      for (const callback of this.subscribers.values()) {
        callback();
      }
    };

    this.ws.onclose = () => {
      this.connected = false;
      this.reconnect();
    };

    this.ws.onerror = (error) => {
      console.error("Websocket error", error);
      this.ws?.close();
    };
  }

  private reconnect() {
    setTimeout(() => {
      this.connect();
    }, 1000);
  }

  subscribe(caller: any, callback: () => void) {
    this.subscribers.set(caller, callback);
  }

  unsubscribe(caller: any) {
    this.subscribers.delete(caller);
  }

  get currentSequenceID(): number | null {
    if (this.queue.length === 0) {
      return null;
    } else {
      return this.queue[this.current_idx];
    }
  }

  get status(): string {
    if (!this.connected) {
      return "Disconnected";
    } else if (this.running) {
      return "Running";
    } else {
      return "Stopped";
    }
  }

  start() {
    axios.post("/v1/player/start");
  }

  stop() {
    axios.post("/v1/player/stop");
  }
}

const playerState = reactive(new PlayerState());
playerState.connect();
export default playerState;
