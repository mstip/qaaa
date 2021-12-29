import { RunEl, MethodEl, UrlEl, RunResultEl } from "./elements.js";

const resultTempalte = `
<div class="col mt-3">
    <hr>
    <h4><strong>Result:</strong>
        {% if result.data.Success %}
        <span class="text-success"><i class="bi-check"></i> Success</span>
        {% else %}
        <span class="text-danger"><i class="bi-x"></i> Error</span>
        {% endif %}
    </h4>
    <div>
        <label class="form-label">Detail</label>
        <input type="text" class="form-control" value="{{result.data.Detail}}" readonly>
    </div>
    <div>
        <label class="form-label">Body</label>
        <pre>{{responseBody}}</pre>
    </div>
</div>
`;

function getChecks() {
  const checks = [];
  document.querySelectorAll(".js-check").forEach((el, i) => {
    checks.push({
      selector: el.querySelector(".js-check-selector").value,
      type: el.querySelector(".js-check-type").value,
      expectedValue: el.querySelector(".js-check-expected-value").value,
    });
  });
  return checks;
}

async function run(event) {
  event.preventDefault();
  try {
    const result = await axios.post("/api/task/run", {
      method: MethodEl.value,
      url: UrlEl.value,
      checks: getChecks(),
    });

    RunResultEl.innerHTML = nunjucks.renderString(resultTempalte, {
      result,
      responseBody: JSON.parse(JSON.stringify(result.data.ResponseBody, 0, 2)),
    });
    console.log(result);
    // RunResultEl.innerHTML = `<hr class="mt-3"><pre>` + JSON.stringify(result, 0, 2) + "</pre>";
  } catch (e) {
    RunResultEl.innerHTML = `<hr class="mt-3"><pre>` + JSON.stringify(e, 0, 2) + "</pre>";
  }
}

export default function Run() {
  RunEl.onclick = run;
}
