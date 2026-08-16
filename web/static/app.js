document.addEventListener('DOMContentLoaded', () => {
  // Global State
  let appState = {
    drivers: [],
    configPath: 'generated.yaml',
    config: {
      driver: 'mysql',
      dsn: '',
      out_dir: 'generated',
      model_out: 'generated/model',
      query_out: 'generated/query',
      validator_out: 'generated/validator',
      validator_package: 'validator',
      tables: [],
      table_prefix: '',
      ignore_fields: ['created_at', 'updated_at', 'deleted_at'],
      tags: ['json', 'form', 'xml', 'url'],
      insert_scene: 'insert',
      update_scene: 'update',
      delete_scene: 'delete',
      type_mappings: {},
      table_configs: {}
    },
    dbTables: [],
    currentEditingTable: ''
  };

  // Init Application
  initApp();

  async function initApp() {
    setupTabNavigation();
    setupTagInputs();
    setupTypeMappingsTable();
    setupTableFieldsEditor();
    setupEventListeners();

    await fetchDrivers();
    await fetchConfig();
  }

  // Navigation Logic
  function setupTabNavigation() {
    const navItems = document.querySelectorAll('.nav-item');
    const tabPanels = document.querySelectorAll('.tab-panel');
    const tabTitle = document.getElementById('tab-title');
    const tabDesc = document.getElementById('tab-desc');

    const tabInfo = {
      db: { title: '🔌 数据库连接与测试', desc: '配置目标数据库驱动及连接字符串 (DSN)，并实时获取数据表列表' },
      paths: { title: '📁 输出路径与包名', desc: '设定模型、Query 操作层和验证结构体的输出目录及 Package 名称' },
      tags: { title: '🏷️ 验证与 Tag 配置', desc: '管理结构体 Tag、验证场景名称以及生成时忽略的数据库列' },
      tables: { title: '📋 数据库表选择', desc: '选择需要生成代码的数据表；留空时默认读取数据库中的全部表' },
      mappings: { title: '⚙️ 全局类型映射', desc: '建立数据库特定数据类型到 Go 语言类型及 Import 包路径的全局映射关系' },
      'table-fields': { title: '🧩 单表字段类型配置', desc: '针对特定表中的列独立指定 Go 类型和 Import 路径，ORM 模型和验证器结构将同步生成' },
      generate: { title: '🚀 执行代码生成', desc: '一键运行生成器，输出 GORM 模型、Query 操作层和 HTTP 校验结构' }
    };

    navItems.forEach(item => {
      item.addEventListener('click', () => {
        const targetTab = item.getAttribute('data-tab');
        
        navItems.forEach(n => n.classList.remove('active'));
        tabPanels.forEach(p => p.classList.remove('active'));

        item.classList.add('active');
        const activePanel = document.getElementById(`panel-${targetTab}`);
        if (activePanel) activePanel.classList.add('active');

        if (tabInfo[targetTab]) {
          tabTitle.textContent = tabInfo[targetTab].title;
          tabDesc.textContent = tabInfo[targetTab].desc;
        }

        // Auto fetch columns if switching to table-fields tab
        if (targetTab === 'table-fields') {
          populateTableFieldSelector();
          const selector = document.getElementById('table-field-selector');
          if (selector.value) {
            fetchTableColumns(selector.value);
          }
        }
      });
    });
  }

  // API Call: Fetch Drivers
  async function fetchDrivers() {
    try {
      const res = await fetch('/api/drivers');
      const data = await res.json();
      if (data.drivers) {
        appState.drivers = data.drivers;
        renderDriverSelect();
        renderPresetsGrid();
      }
    } catch (err) {
      showToast('获取数据库驱动列表失败: ' + err.message, 'error');
    }
  }

  function renderDriverSelect() {
    const select = document.getElementById('driver-select');
    select.innerHTML = '';
    appState.drivers.forEach(d => {
      const opt = document.createElement('option');
      opt.value = d.name;
      opt.textContent = `${d.label} (${d.name})`;
      select.appendChild(opt);
    });

    select.addEventListener('change', (e) => {
      const driver = appState.drivers.find(d => d.name === e.target.value);
      if (driver) {
        document.getElementById('dsn-hint').textContent = `示例: ${driver.example_dsn}`;
      }
    });
  }

  function renderPresetsGrid() {
    const grid = document.getElementById('presets-grid');
    grid.innerHTML = '';
    appState.drivers.forEach(d => {
      const card = document.createElement('div');
      card.className = 'preset-card';
      card.innerHTML = `
        <h4>${d.label}</h4>
        <p>${d.description}</p>
        <code>${d.example_dsn}</code>
      `;
      card.addEventListener('click', () => {
        document.getElementById('driver-select').value = d.name;
        document.getElementById('dsn-input').value = d.example_dsn;
        document.getElementById('dsn-hint').textContent = `示例: ${d.example_dsn}`;
        showToast(`已应用 ${d.label} 示例 DSN`, 'info');
      });
      grid.appendChild(card);
    });
  }

  // API Call: Fetch Config
  async function fetchConfig() {
    try {
      const res = await fetch('/api/config');
      const data = await res.json();
      if (data.config) {
        appState.config = data.config;
        if (!appState.config.table_configs) {
          appState.config.table_configs = {};
        }
        appState.configPath = data.path || 'generated.yaml';
        document.getElementById('active-config-path').textContent = appState.configPath;
        bindConfigToUI();

        // Auto test connection if DSN exists
        if (appState.config.dsn) {
          testDBConnection(true);
        }
      }
    } catch (err) {
      showToast('加载配置文件失败: ' + err.message, 'error');
    }
  }

  function bindConfigToUI() {
    const cfg = appState.config;
    if (cfg.driver) document.getElementById('driver-select').value = cfg.driver;
    if (cfg.dsn) document.getElementById('dsn-input').value = cfg.dsn;
    if (cfg.out_dir) document.getElementById('out-dir').value = cfg.out_dir;
    if (cfg.model_out) document.getElementById('model-out').value = cfg.model_out;
    if (cfg.query_out) document.getElementById('query-out').value = cfg.query_out;
    if (cfg.validator_out) document.getElementById('validator-out').value = cfg.validator_out;
    if (cfg.validator_package) document.getElementById('validator-pkg').value = cfg.validator_package;
    if (cfg.table_prefix) document.getElementById('table-prefix').value = cfg.table_prefix;
    if (cfg.insert_scene) document.getElementById('insert-scene').value = cfg.insert_scene;
    if (cfg.update_scene) document.getElementById('update-scene').value = cfg.update_scene;
    if (cfg.delete_scene) document.getElementById('delete-scene').value = cfg.delete_scene;

    renderTagPills('tags-wrapper', cfg.tags || []);
    renderTagPills('ignore-wrapper', cfg.ignore_fields || []);
    renderTypeMappingsRows(cfg.type_mappings || {});
    updateSelectedTableCount();
    populateTableFieldSelector();
  }

  function collectConfigFromUI() {
    saveCurrentTableFieldsToState();

    return {
      driver: document.getElementById('driver-select').value,
      dsn: document.getElementById('dsn-input').value,
      out_dir: document.getElementById('out-dir').value,
      model_out: document.getElementById('model-out').value,
      query_out: document.getElementById('query-out').value,
      validator_out: document.getElementById('validator-out').value,
      validator_package: document.getElementById('validator-pkg').value,
      table_prefix: document.getElementById('table-prefix').value,
      insert_scene: document.getElementById('insert-scene').value,
      update_scene: document.getElementById('update-scene').value,
      delete_scene: document.getElementById('delete-scene').value,
      tags: collectTagPills('tags-wrapper'),
      ignore_fields: collectTagPills('ignore-wrapper'),
      tables: getSelectedTables(),
      type_mappings: collectTypeMappingsFromUI(),
      table_configs: appState.config.table_configs || {}
    };
  }

  // Tag Inputs Helper
  function setupTagInputs() {
    setupSingleTagInput('tags-wrapper', 'tag-input-field');
    setupSingleTagInput('ignore-wrapper', 'ignore-input-field');
  }

  function setupSingleTagInput(wrapperId, inputId) {
    const wrapper = document.getElementById(wrapperId);
    const input = document.getElementById(inputId);

    input.addEventListener('keydown', (e) => {
      if (e.key === 'Enter' || e.key === ',') {
        e.preventDefault();
        const val = input.value.trim().replace(/,/g, '');
        if (val) {
          addTagPill(wrapperId, val);
          input.value = '';
        }
      }
    });
  }

  function addTagPill(wrapperId, text) {
    const wrapper = document.getElementById(wrapperId);
    const input = wrapper.querySelector('input');
    const existing = collectTagPills(wrapperId);
    if (existing.includes(text)) return;

    const pill = document.createElement('span');
    pill.className = 'tag-pill';
    pill.dataset.value = text;
    pill.innerHTML = `${text} <span class="remove-btn">&times;</span>`;
    
    pill.querySelector('.remove-btn').addEventListener('click', () => {
      pill.remove();
    });

    wrapper.insertBefore(pill, input);
  }

  function renderTagPills(wrapperId, items) {
    const wrapper = document.getElementById(wrapperId);
    const input = wrapper.querySelector('input');
    wrapper.querySelectorAll('.tag-pill').forEach(el => el.remove());

    items.forEach(item => {
      addTagPill(wrapperId, item);
    });
  }

  function collectTagPills(wrapperId) {
    const wrapper = document.getElementById(wrapperId);
    const pills = wrapper.querySelectorAll('.tag-pill');
    return Array.from(pills).map(p => p.dataset.value);
  }

  // Custom Type Mappings Table Logic
  function setupTypeMappingsTable() {
    document.getElementById('btn-add-mapping').addEventListener('click', () => {
      addTypeMappingRow('', '', '');
    });

    document.getElementById('preset-uuid').addEventListener('click', () => {
      addTypeMappingRow('uuid', 'uuid.UUID', 'github.com/google/uuid');
    });

    document.getElementById('preset-jsonb').addEventListener('click', () => {
      addTypeMappingRow('jsonb', 'datatypes.JSON', 'gorm.io/datatypes');
    });
  }

  function addTypeMappingRow(dbType, goType, importPath) {
    const tbody = document.getElementById('mappings-body');
    const tr = document.createElement('tr');
    tr.innerHTML = `
      <td><input type="text" class="mapping-db" value="${dbType}" placeholder="例如: uuid, jsonb"></td>
      <td><input type="text" class="mapping-go" value="${goType}" placeholder="例如: uuid.UUID"></td>
      <td><input type="text" class="mapping-import" value="${importPath}" placeholder="例如: github.com/google/uuid"></td>
      <td><button class="btn btn-xs btn-outline btn-del-row" style="color:var(--accent-red);border-color:var(--accent-red);">删除</button></td>
    `;
    tr.querySelector('.btn-del-row').addEventListener('click', () => tr.remove());
    tbody.appendChild(tr);
  }

  function renderTypeMappingsRows(mappings) {
    const tbody = document.getElementById('mappings-body');
    tbody.innerHTML = '';
    Object.entries(mappings).forEach(([key, val]) => {
      addTypeMappingRow(val.db_type || key, val.go_type || '', val.import_path || '');
    });
  }

  function collectTypeMappingsFromUI() {
    const mappings = {};
    const rows = document.querySelectorAll('#mappings-body tr');
    rows.forEach(tr => {
      const dbType = tr.querySelector('.mapping-db').value.trim();
      const goType = tr.querySelector('.mapping-go').value.trim();
      const importPath = tr.querySelector('.mapping-import').value.trim();
      if (dbType && goType) {
        mappings[dbType] = {
          db_type: dbType,
          go_type: goType,
          import_path: importPath
        };
      }
    });
    return mappings;
  }

  // Database Connection Test
  async function testDBConnection(silent = false) {
    const driver = document.getElementById('driver-select').value;
    const dsn = document.getElementById('dsn-input').value.trim();
    const resultBox = document.getElementById('db-test-result');

    if (!driver || !dsn) {
      if (!silent) showToast('请填写驱动和 DSN 连接字符串', 'error');
      return;
    }

    if (!silent) {
      resultBox.style.display = 'block';
      resultBox.className = 'db-test-result';
      resultBox.textContent = '⏳ 正在测试数据库连接并获取表...';
    }

    try {
      const res = await fetch('/api/db/test', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ driver, dsn })
      });
      const data = await res.json();

      if (data.success) {
        appState.dbTables = data.tables || [];
        renderTablesGrid();
        populateTableFieldSelector();
        
        resultBox.style.display = 'block';
        resultBox.className = 'db-test-result success';
        resultBox.textContent = `✅ 连接成功！数据库中共发现 ${data.table_count} 张表。`;
        if (!silent) showToast(`数据库连接成功，共 ${data.table_count} 张表`, 'success');
      } else {
        resultBox.style.display = 'block';
        resultBox.className = 'db-test-result error';
        resultBox.textContent = `❌ 连接失败: ${data.error}`;
        if (!silent) showToast(`数据库连接失败: ${data.error}`, 'error');
      }
    } catch (err) {
      resultBox.style.display = 'block';
      resultBox.className = 'db-test-result error';
      resultBox.textContent = `❌ 请求错误: ${err.message}`;
      if (!silent) showToast(`请求失败: ${err.message}`, 'error');
    }
  }

  // Render Table Selector Grid
  function renderTablesGrid() {
    const grid = document.getElementById('tables-grid');
    const selectedTables = appState.config.tables || [];

    if (appState.dbTables.length === 0) {
      grid.innerHTML = '<div class="empty-state">💡 数据库中未读取到任何表。</div>';
      return;
    }

    grid.innerHTML = '';
    appState.dbTables.forEach(tableName => {
      const isChecked = selectedTables.length === 0 || selectedTables.includes(tableName);
      const item = document.createElement('div');
      item.className = `table-badge-item ${isChecked ? 'selected' : ''}`;
      item.dataset.table = tableName;
      item.innerHTML = `
        <input type="checkbox" ${isChecked ? 'checked' : ''}>
        <span>${tableName}</span>
      `;

      item.addEventListener('click', (e) => {
        if (e.target.tagName !== 'INPUT') {
          const chk = item.querySelector('input');
          chk.checked = !chk.checked;
        }
        const checked = item.querySelector('input').checked;
        item.classList.toggle('selected', checked);
        updateSelectedTableCount();
      });

      grid.appendChild(item);
    });

    updateSelectedTableCount();
  }

  function getSelectedTables() {
    const selected = [];
    document.querySelectorAll('.table-badge-item').forEach(item => {
      if (item.querySelector('input').checked) {
        selected.push(item.dataset.table);
      }
    });

    if (selected.length === appState.dbTables.length) {
      return [];
    }
    return selected;
  }

  function updateSelectedTableCount() {
    const selected = getSelectedTables();
    const badge = document.getElementById('selected-table-count');
    if (selected.length === 0) {
      badge.textContent = '全部 (' + (appState.dbTables.length || 0) + ')';
    } else {
      badge.textContent = `${selected.length} / ${appState.dbTables.length}`;
    }
  }

  // Table Fields Editor Logic
  function setupTableFieldsEditor() {
    const selector = document.getElementById('table-field-selector');
    const btnFetch = document.getElementById('btn-fetch-table-fields');

    selector.addEventListener('change', (e) => {
      saveCurrentTableFieldsToState();
      const tableName = e.target.value;
      if (tableName) {
        fetchTableColumns(tableName);
      } else {
        document.getElementById('table-fields-editor-container').innerHTML = 
          '<div class="empty-state">💡 请在右上角下拉列表中选择需要单独配置字段类型的数据库表。</div>';
      }
    });

    btnFetch.addEventListener('click', () => {
      const tableName = selector.value;
      if (tableName) {
        fetchTableColumns(tableName);
      } else {
        showToast('请先选择要载入字段的数据库表', 'error');
      }
    });
  }

  function populateTableFieldSelector() {
    const selector = document.getElementById('table-field-selector');
    const currentVal = selector.value;
    selector.innerHTML = '<option value="">-- 选择数据表 --</option>';

    appState.dbTables.forEach(tableName => {
      const opt = document.createElement('option');
      opt.value = tableName;
      opt.textContent = tableName;
      selector.appendChild(opt);
    });

    if (currentVal && appState.dbTables.includes(currentVal)) {
      selector.value = currentVal;
    }
  }

  async function fetchTableColumns(tableName) {
    const container = document.getElementById('table-fields-editor-container');
    container.innerHTML = `<div class="empty-state">⏳ 正在读取表 [${tableName}] 的列结构与元数据...</div>`;
    appState.currentEditingTable = tableName;

    const driver = document.getElementById('driver-select').value;
    const dsn = document.getElementById('dsn-input').value.trim();

    if (!driver || !dsn) {
      container.innerHTML = '<div class="db-test-result error">❌ 请先在“数据库连接”页面填写 Driver 和 DSN。</div>';
      return;
    }

    try {
      const res = await fetch('/api/db/columns', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          driver,
          dsn,
          table: tableName,
          config: collectConfigFromUI()
        })
      });
      const data = await res.json();

      if (data.success && data.columns) {
        renderTableFieldsEditor(tableName, data.columns);
      } else {
        container.innerHTML = `<div class="db-test-result error">❌ 读取表结构失败: ${data.error || '未知错误'}</div>`;
      }
    } catch (err) {
      container.innerHTML = `<div class="db-test-result error">❌ 请求失败: ${err.message}</div>`;
    }
  }

  function renderTableFieldsEditor(tableName, columns) {
    const container = document.getElementById('table-fields-editor-container');
    const existingTableConfig = (appState.config.table_configs && appState.config.table_configs[tableName]) || { fields: {} };
    const existingFields = existingTableConfig.fields || {};

    let html = `
      <div class="table-responsive">
        <table class="data-table">
          <thead>
            <tr>
              <th>数据库列名</th>
              <th>DB 数据类型</th>
              <th>默认推导 Go 类型</th>
              <th>自定义 Go 类型 (GoType)</th>
              <th>Import 路径 (ImportPath)</th>
              <th>快捷预设</th>
            </tr>
          </thead>
          <tbody id="table-fields-body">
    `;

    columns.forEach(col => {
      const savedOverride = existingFields[col.name] || { go_type: '', import_path: '' };
      const customGoType = savedOverride.go_type || '';
      const customImport = savedOverride.import_path || '';

      html += `
        <tr data-column="${col.name}">
          <td>
            <strong>${col.name}</strong>
            ${col.primary ? '<span class="badge" style="margin-left:4px;background:rgba(210,153,34,0.2);color:#d29922;">PK</span>' : ''}
            ${col.comment ? `<div style="font-size:11px;color:var(--text-muted);">${col.comment}</div>` : ''}
          </td>
          <td><code style="font-size:12px;">${col.database_type}</code></td>
          <td><span style="color:var(--text-muted);font-family:var(--font-mono);font-size:12px;">${col.go_type}</span></td>
          <td><input type="text" class="field-custom-gotype" value="${customGoType}" placeholder="例如: custom.MyType"></td>
          <td><input type="text" class="field-custom-import" value="${customImport}" placeholder="例如: myproject/custom"></td>
          <td>
            <div class="flex-gap" style="gap:4px;">
              <button class="btn btn-xs btn-outline btn-preset-json" data-col="${col.name}">JSON</button>
              <button class="btn btn-xs btn-outline btn-preset-uuid" data-col="${col.name}">UUID</button>
              <button class="btn btn-xs btn-outline btn-preset-clear" data-col="${col.name}" style="color:var(--text-muted);border-color:var(--border-color);">清空</button>
            </div>
          </td>
        </tr>
      `;
    });

    html += `
          </tbody>
        </table>
      </div>
    `;

    container.innerHTML = html;

    // Attach preset button handlers
    container.querySelectorAll('.btn-preset-json').forEach(btn => {
      btn.addEventListener('click', () => {
        const tr = btn.closest('tr');
        tr.querySelector('.field-custom-gotype').value = 'datatypes.JSON';
        tr.querySelector('.field-custom-import').value = 'gorm.io/datatypes';
        saveCurrentTableFieldsToState();
      });
    });

    container.querySelectorAll('.btn-preset-uuid').forEach(btn => {
      btn.addEventListener('click', () => {
        const tr = btn.closest('tr');
        tr.querySelector('.field-custom-gotype').value = 'uuid.UUID';
        tr.querySelector('.field-custom-import').value = 'github.com/google/uuid';
        saveCurrentTableFieldsToState();
      });
    });

    container.querySelectorAll('.btn-preset-clear').forEach(btn => {
      btn.addEventListener('click', () => {
        const tr = btn.closest('tr');
        tr.querySelector('.field-custom-gotype').value = '';
        tr.querySelector('.field-custom-import').value = '';
        saveCurrentTableFieldsToState();
      });
    });

    // Auto save on input change
    container.querySelectorAll('input').forEach(input => {
      input.addEventListener('change', saveCurrentTableFieldsToState);
    });
  }

  function saveCurrentTableFieldsToState() {
    const tableName = appState.currentEditingTable;
    if (!tableName) return;

    const rows = document.querySelectorAll('#table-fields-body tr');
    if (!rows || rows.length === 0) return;

    if (!appState.config.table_configs) {
      appState.config.table_configs = {};
    }

    const fieldsMap = {};
    rows.forEach(tr => {
      const colName = tr.dataset.column;
      const goType = tr.querySelector('.field-custom-gotype').value.trim();
      const importPath = tr.querySelector('.field-custom-import').value.trim();

      if (goType) {
        fieldsMap[colName] = {
          go_type: goType,
          import_path: importPath
        };
      }
    });

    if (Object.keys(fieldsMap).length > 0) {
      appState.config.table_configs[tableName] = { fields: fieldsMap };
    } else {
      delete appState.config.table_configs[tableName];
    }
  }

  // Setup General Event Listeners
  function setupEventListeners() {
    document.getElementById('btn-test-db').addEventListener('click', () => testDBConnection(false));

    document.getElementById('btn-save-config').addEventListener('click', saveConfig);
    document.getElementById('btn-load-config').addEventListener('click', fetchConfig);

    document.getElementById('btn-top-generate').addEventListener('click', () => {
      document.querySelector('[data-tab="generate"]').click();
      runGenerate();
    });

    document.getElementById('btn-run-generate').addEventListener('click', runGenerate);

    document.getElementById('btn-select-all-tables').addEventListener('click', () => {
      document.querySelectorAll('.table-badge-item').forEach(item => {
        item.querySelector('input').checked = true;
        item.classList.add('selected');
      });
      updateSelectedTableCount();
    });

    document.getElementById('btn-clear-tables').addEventListener('click', () => {
      document.querySelectorAll('.table-badge-item').forEach(item => {
        item.querySelector('input').checked = false;
        item.classList.remove('selected');
      });
      updateSelectedTableCount();
    });

    document.getElementById('table-search-input').addEventListener('input', (e) => {
      const q = e.target.value.toLowerCase();
      document.querySelectorAll('.table-badge-item').forEach(item => {
        const text = item.dataset.table.toLowerCase();
        item.style.display = text.includes(q) ? 'flex' : 'none';
      });
    });

    document.getElementById('btn-clear-log').addEventListener('click', () => {
      document.getElementById('console-output').textContent = '';
    });
  }

  // Save Config API Call
  async function saveConfig() {
    const cfg = collectConfigFromUI();
    const path = appState.configPath || 'generated.yaml';

    try {
      const res = await fetch('/api/config/save', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ path, config: cfg })
      });
      const data = await res.json();

      if (data.success) {
        showToast(`配置已成功保存至 ${path}`, 'success');
      } else {
        showToast(`保存失败: ${data.error}`, 'error');
      }
    } catch (err) {
      showToast(`请求错误: ${err.message}`, 'error');
    }
  }

  // Code Generation API Call
  async function runGenerate() {
    const cfg = collectConfigFromUI();
    const autoSave = document.getElementById('chk-auto-save').checked;
    const savePath = appState.configPath || 'generated.yaml';

    const outputEl = document.getElementById('console-output');
    const resultFilesCard = document.getElementById('result-files-card');
    const filesList = document.getElementById('generated-files-list');

    outputEl.textContent = `🚀 [${new Date().toLocaleTimeString()}] 开始代码生成任务...\n`;
    outputEl.textContent += `⚙️ 数据库驱动: ${cfg.driver}\n`;
    outputEl.textContent += `📁 输出根路径: ${cfg.out_dir}\n`;
    if (cfg.table_configs && Object.keys(cfg.table_configs).length > 0) {
      outputEl.textContent += `🧩 已应用单表字段类型配置: ${Object.keys(cfg.table_configs).join(', ')}\n`;
    }
    outputEl.textContent += `⏳ 正在连接数据库并构建模型与验证文件...\n`;

    resultFilesCard.style.display = 'none';

    try {
      const res = await fetch('/api/generate', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          config: cfg,
          save_path: savePath,
          auto_save: autoSave
        })
      });
      const data = await res.json();

      if (data.success) {
        outputEl.textContent += `✅ [${new Date().toLocaleTimeString()}] 代码生成完毕！\n`;
        outputEl.textContent += `📦 生成路径模型: ${cfg.model_out}\n`;
        outputEl.textContent += `📦 生成路径 Query: ${cfg.query_out}\n`;
        outputEl.textContent += `📦 生成路径 Validator: ${cfg.validator_out}\n`;

        showToast('代码生成成功！', 'success');

        if (data.files && data.files.length > 0) {
          resultFilesCard.style.display = 'block';
          filesList.innerHTML = '';
          data.files.forEach(f => {
            const li = document.createElement('li');
            li.textContent = f;
            filesList.appendChild(li);
          });
        }
      } else {
        outputEl.textContent += `❌ [${new Date().toLocaleTimeString()}] 生成失败!\n错误详情: ${data.error}\n`;
        showToast('生成失败: ' + data.error, 'error');
      }
    } catch (err) {
      outputEl.textContent += `❌ [${new Date().toLocaleTimeString()}] 网络/服务器请求异常: ${err.message}\n`;
      showToast('请求异常: ' + err.message, 'error');
    }
  }

  // Toast Notification Helper
  function showToast(msg, type = 'info') {
    const container = document.getElementById('toast-container');
    const toast = document.createElement('div');
    toast.className = `toast ${type}`;
    toast.textContent = msg;

    container.appendChild(toast);
    setTimeout(() => {
      toast.style.opacity = '0';
      setTimeout(() => toast.remove(), 300);
    }, 3500);
  }
});
