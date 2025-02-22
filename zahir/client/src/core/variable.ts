type Value = number | string | boolean;

export default class Variable {
  id!: number;
  name!: string;
  description!: string;
  defaultValue!: Value;

  static fromResource(resource: any): Variable {
    const variable = new Variable();
    variable.id = resource.id;
    variable.name = resource.name;
    variable.description = resource.description;
    variable.defaultValue = resource.default_value;
    return variable;
  }
}
