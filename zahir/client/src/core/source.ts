import axios from "axios";
import Variable from "./variable";

export default class Source {
  id!: number;
  name!: string;
  description!: string;
  type!: string;

  private variables!: Variable[];

  static fromResource(resource: any): Source {
    const source = new Source();
    source.id = resource.id;
    source.name = resource.name;
    source.description = resource.description;
    source.type = resource.type;
    source.variables = resource.variables?.map(Variable.fromResource);
    return source;
  }

  static async fetchAll(): Promise<Source[]> {
    const response = await axios.get("/v1/sources");
    return response.data.map(Source.fromResource);
  }

  static async fetchById(id: number): Promise<Source> {
    const response = await axios.get(`/v1/sources/${id}`);
    return Source.fromResource(response.data);
  }

  static async createNew() {
    await axios.post("/v1/sources");
  }

  getVariables(): Variable[] {
    if (this.variables === undefined) {
      throw new Error("Variables not loaded");
    }
    return this.variables;
  }

  removeVariable(variableId: number) {
    this.variables = this.variables.filter((variable) => variable.id !== variableId);
  }
}
