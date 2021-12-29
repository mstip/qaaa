import { ChecksEl } from "./elements.js";

const checkTemplate = `
   <button class="js-addNewCheck btn btn-primary">Add new check</button>
   <form>
    {% for check in checksData %}
    <div class="js-check js-check-{{loop.index}}">
      <div>
        <label class="form-label">Selector</label>
        <input type="text" class="form-control js-check-selector" name="selector" value="{{check.Selector }}" />
      </div>
      <div>
        <label class="form-label">Type</label>
        <select name="type" class="form-control js-check-type">
          <option value="equals" {% if check.Type === "equals" %}selected{%endif%} >Equals</option>
          <option value="equals_not" {% if check.Type === "equalsNot"%}selected{%endif%}>Equals Not</option>
          <option value="contains" {% if check.Type === "contains"%}selected{%endif%}>Contains</option>
          <option value="contains_not" {% if check.Type === "containsNot"%}selected{%endif%}>Contains Not</option>
          <option value="count" {% if check.Type === "count"%}selected{%endif%}>Count</option>
        </select>
      </div>
      <div>
        <label class="form-label">Expected Value</label>
        <input type="text" class="form-control js-check-expected-value" name="expectedValue" value="{{check.Value }}" />
      </div>
    </div>
    <hr>
    {% endfor %}
  </form>
`;

function render(checksData) {
  ChecksEl.innerHTML = nunjucks.renderString(checkTemplate, { checksData });

  document.querySelector(".js-addNewCheck").onclick = () => {
    checksData.push({ Selector: "", Type: "", Value: "" });
    render(checksData);
  };
}

export default async function Checks() {
  try {
    const checksRequest = await axios.get("/api/task/detail/0/checks");
    render(checksRequest.data);
  } catch (error) {
    console.log(error);
  }
}
