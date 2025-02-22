import Variable from "./variable";
import Sequence from "./sequence";

export default class GraphCollection {
  private graphs: Graph[] = [];

  public sequence: Sequence;
  public popover: Popover;

  constructor(sequence: Sequence) {
    this.sequence = sequence;
    this.popover = new Popover(sequence);
  }

  addGraph(variable: Variable) {
    const svg = document.getElementById(`svg-${variable.id}`) as SVGSVGElement | null;
    const title = document.getElementById(`title-${variable.id}`) as HTMLElement | null;
    if (svg && title) {
      this.graphs.push(new Graph(this, variable, svg, title));
    } else {
      console.error(`Could not find svg or title for variable ${variable.id}`);
    }
  }

  unHighlightAll() {
    this.graphs.forEach((graph) => graph.unHighligh());
  }

  highlightAll(x: number) {
    this.graphs.forEach((graph) => graph.highlight(x));
  }

  setAllTitlesColor(color: string) {
    this.graphs.forEach((graph) => graph.setTitleColor(color));
  }

  placePlayLine(x: number | null) {
    this.graphs.forEach((graph) => graph.placePlayLine(x));
  }

  destroy() {}
}

class Popover {
  private selectedGraph!: Graph;
  private selectedX!: number;

  constructor(private sequence: Sequence) {
    this.submitOneButton.addEventListener("click", () => this.submit(true));
    this.submitAllButton.addEventListener("click", () => this.submit(false));
    this.input.addEventListener("keyup", (e) => {
      if (e.key === "Enter") this.submit();
    });
  }

  private get element(): HTMLElement {
    return document.getElementById("popover")!;
  }

  private get title(): HTMLElement {
    return document.getElementById("popover-title")!;
  }

  private get input(): HTMLInputElement {
    return document.getElementById("popover-input") as HTMLInputElement;
  }

  private get submitOneButton(): HTMLElement {
    return document.getElementById("popover-button-one")!;
  }

  private get submitAllButton(): HTMLElement {
    return document.getElementById("popover-button-all")!;
  }

  private get heightPx(): number {
    return this.element.getBoundingClientRect().height;
  }

  private get widthPx(): number {
    return this.element.getBoundingClientRect().width;
  }

  private async submit(one: boolean = true) {
    const variableId = this.selectedGraph.getVariableId();
    const val = parseFloat(this.input.value);
    const arr = this.selectedGraph.getValues();

    if (one) {
      arr[this.selectedX] = parseFloat(this.input.value);
    } else {
      for (let i = this.selectedX; i < arr.length; i++) {
        arr[i] = val;
      }
    }

    await this.sequence.setValues(variableId, arr);
    this.selectedGraph.draw();
    this.hide();
  }

  show(top: number, left: number, graph: Graph, x: number) {
    this.selectedGraph = graph;
    this.selectedX = x;

    this.element.style.display = "block";
    // this.title.innerText = graph.getVariableName();
    this.input.value = graph.getValues()[x].toString();
    this.input.focus();

    this.element.style.top = `${top - this.heightPx}px`;
    this.element.style.left = `${left - this.widthPx / 2}px`;
  }

  hide() {
    this.element.style.display = "none";
  }

  isVisible(): boolean {
    return this.element.style.display === "block";
  }
}

class Graph {
  private readonly ns = "http://www.w3.org/2000/svg";
  private readonly highlightColor: string = "#cccccc";
  private readonly playLineColor: string = "#179654";
  private readonly heightPx: number = 70;

  private refLines!: Element[];
  private playLine!: Element;
  private height!: number;
  private showSquareLines: boolean = false;

  constructor(
    private collection: GraphCollection,
    private variable: Variable,
    private svg: SVGSVGElement,
    private title: HTMLElement
  ) {
    this.draw();
  }

  private get length(): number {
    return this.collection.sequence.length;
  }

  getVariableName(): string {
    return this.variable.name;
  }

  getVariableId(): number {
    return this.variable.id;
  }

  getValues(): number[] {
    return this.collection.sequence.getValues(this.variable.id);
  }

  calculateHeight() {
    this.height = this.heightPx / (this.svg.getBoundingClientRect().width / this.length);
  }

  private addLine(x1: number, y1: number, x2: number, y2: number, color: string): SVGElement {
    const line = document.createElementNS(this.ns, "line");
    line.setAttribute("x1", x1.toString());
    line.setAttribute("y1", y1.toString());
    line.setAttribute("x2", x2.toString());
    line.setAttribute("y2", y2.toString());
    line.setAttribute("stroke", color);
    line.setAttribute("stroke-width", (this.height / 50).toString());
    this.svg.appendChild(line);
    return line;
  }

  draw() {
    this.svg.innerHTML = "";
    this.svg.setAttribute("width", "100%");
    this.svg.setAttribute("height", `${this.heightPx}px`);
    this.calculateHeight();

    this.svg.setAttribute("viewBox", `0 0 ${this.length} ${this.height}`);
    this.svg.setAttribute("preserveAspectRatio", "none");
    this.svg.addEventListener("resize", () => this.draw());

    this.drawBlocksAndLines();
    this.drawPlayLine();
    this.drawValues();
  }

  drawBlocksAndLines() {
    this.refLines = [];

    for (let x = 0; x < this.length; x++) {
      const line = this.addLine(x + 0.5, 0, x + 0.5, this.height, "transparent");

      const rect = document.createElementNS(this.ns, "rect");
      rect.setAttribute("x", x.toString());
      rect.setAttribute("y", "0");
      rect.setAttribute("width", "1");
      rect.setAttribute("height", this.height.toString());
      rect.setAttribute("fill", "transparent");

      const block = document.createElementNS(this.ns, "g");
      block.appendChild(rect);
      block.appendChild(line);
      block.setAttribute("cursor", "pointer");
      block.addEventListener("mouseenter", () => this.onMouseEnter(x));
      block.addEventListener("mouseleave", () => this.onMouseLeave());
      block.addEventListener("mouseup", () => this.onMouseUp(x));

      this.refLines.push(line);
      this.svg.appendChild(block);
    }
  }

  drawPlayLine() {
    this.playLine = this.addLine(0, 0, 0, this.height, "transparent");
  }

  drawValues() {
    const values = this.getValues();

    const min = Math.min(...values);
    const max = Math.max(...values);
    const r = 0.15;
    const h = this.height;
    const scale = (v: number) => {
      return max === min ? 0.5 * h : h * (r + (1 - 2 * r) * ((max - v) / (max - min)));
    };

    for (let x = 0; x < this.length - 1; x++) {
      const v0 = values[x];
      const v1 = values[x + 1];
      const x0 = x + 0.5;
      const x1 = x + 1.5;
      const y0 = scale(v0);
      const y1 = scale(v1);

      if (this.showSquareLines) {
        this.addLine(x0, y0, x1, y0, "#0d6efd");
        this.addLine(x1, y0, x1, y1, "#0d6efd");
      } else {
        this.addLine(x0, y0, x1, y1, "#0d6efd");
      }
    }
  }

  onMouseEnter(x: number) {
    this.collection.unHighlightAll();
    this.collection.highlightAll(x);
    this.collection.setAllTitlesColor("black");
    this.setTitleColor("#0d6efd");
  }

  onMouseLeave() {
    this.collection.unHighlightAll();
    this.collection.setAllTitlesColor("black");
  }

  onMouseUp(x: number) {
    const r = this.refLines[x].getBoundingClientRect();
    this.collection.popover.show(r.top, r.left, this, x);
    this.draw();
  }

  unHighligh() {
    this.refLines.forEach((line) => line.setAttribute("stroke", "transparent"));
  }

  highlight(x: number) {
    this.unHighligh();
    this.refLines[x].setAttribute("stroke", this.highlightColor);
  }

  setBackgroundColor(color: string) {
    this.svg.style.backgroundColor = color;
  }

  setTitleColor(color: string) {
    this.title.style.color = color;
  }

  placePlayLine(x: number | null) {
    if (x === null) {
      this.playLine.setAttribute("stroke", "transparent");
      return;
    } else {
      this.playLine.setAttribute("x1", (x + 0.5).toString());
      this.playLine.setAttribute("x2", (x + 0.5).toString());
      this.playLine.setAttribute("stroke", this.playLineColor);
    }
  }
}
