import axios from "axios";
import Variable from "./variable";
import Source from "./source";

interface VariableValues {
  variable_id: number;
  sequence_id: number;
  values: number[];
}

export default class Sequence {
  id!: number;
  name!: string;
  description!: string;
  length!: number;

  // This will only be available after calling fetchById
  private sources!: Source[];
  private values!: Map<number, number[]>; // VariableID -> Values

  static fromResource(resource: any): Sequence {
    const sequence = new Sequence();
    sequence.id = resource.id;
    sequence.name = resource.name;
    sequence.description = resource.description;
    sequence.length = resource.length;
    return sequence;
  }

  static async fetchAll(): Promise<Sequence[]> {
    const response = await axios.get("/v1/sequences");
    return response.data.map(Sequence.fromResource);
  }

  static async fetchById(id: number): Promise<Sequence> {
    const response = await axios.get(`/v1/sequences/${id}`);
    const sequence = Sequence.fromResource(response.data.sequence);
    sequence.sources = response.data.sources.map(Source.fromResource);
    sequence.values = new Map<number, number[]>(
      response.data.values.map((v: VariableValues) => [v.variable_id, v.values])
    );
    return sequence;
  }

  async run() {
    const response = await axios.post(`/v1/sequences/${this.id}/run`);
    return response.data;
  }

  getSources(): Source[] {
    if (this.sources === undefined) {
      throw new Error("Sources not loaded");
    }
    return this.sources;
  }

  getVariables(): Variable[] {
    const variables = [];
    for (const source of this.getSources()) {
      variables.push(...source.getVariables());
    }
    return variables;
  }

  getValues(variableId: number): number[] {
    if (this.values === undefined) {
      throw new Error("Values not loaded");
    }
    return this.values.get(variableId) || [];
  }

  async setValues(variableId: number, values: number[]) {
    if (this.values === undefined) {
      throw new Error("Values not loaded");
    }
    const response = await axios.post(`/v1/sequences/${this.id}/values`, {
      variable_id: variableId,
      values: values,
    });
    if (response.status !== 200) {
      throw new Error("Failed to set values");
    }
    this.values.set(variableId, values);
  }
}
