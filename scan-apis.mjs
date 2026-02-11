#!/usr/bin/env node

import { readFileSync, readdirSync, statSync } from 'fs';
import { join } from 'path';

const apiEndpoints = new Set();
const filesByEndpoint = new Map();

function scanDirectory(dir, extensions = ['.jsx', '.tsx', '.js', '.ts', '.astro']) {
  try {
    const files = readdirSync(dir);
    
    for (const file of files) {
      const fullPath = join(dir, file);
      const stat = statSync(fullPath);
      
      if (stat.isDirectory()) {
        if (!file.includes('node_modules') && !file.includes('.git')) {
          scanDirectory(fullPath, extensions);
        }
      } else if (extensions.some(ext => file.endsWith(ext))) {
        scanFile(fullPath);
      }
    }
  } catch (error) {
    console.error(`Error scanning ${dir}:`, error.message);
  }
}

function scanFile(filePath) {
  try {
    const content = readFileSync(filePath, 'utf-8');
    
    // Regex patterns para encontrar APIs
    const patterns = [
      /fetch\s*\(\s*['"`]([^'"`]*\/api\/[^'"`]*?)['"`]/g,
      /fetch\s*\(\s*`([^`]*\/api\/[^`]*?)`/g,
      /axios\.[a-z]+\s*\(\s*['"`]([^'"`]*\/api\/[^'"`]*?)['"`]/g,
      /url:\s*['"`]([^'"`]*\/api\/[^'"`]*?)['"`]/g,
      /action=['"`]([^'"`]*\/api\/[^'"`]*?)['"`]/g,
      /href=['"`]([^'"`]*\/api\/[^'"`]*?)['"`]/g,
    ];
    
    for (const pattern of patterns) {
      let match;
      while ((match = pattern.exec(content)) !== null) {
        let endpoint = match[1];
        
        // Limpiar query params y variables
        endpoint = endpoint.split('?')[0];
        endpoint = endpoint.replace(/\$\{[^}]+\}/g, ':param');
        endpoint = endpoint.replace(/\+[^+]*\+/g, '');
        endpoint = endpoint.trim();
        
        if (endpoint.includes('/api/')) {
          apiEndpoints.add(endpoint);
          
          if (!filesByEndpoint.has(endpoint)) {
            filesByEndpoint.set(endpoint, []);
          }
          filesByEndpoint.get(endpoint).push(filePath);
        }
      }
    }
  } catch (error) {
    // Ignorar errores de lectura
  }
}

console.log('🔍 Escaneando APIs en uso...\n');

// Escanear directorios
scanDirectory('./app/src');
scanDirectory('./caja/src');
scanDirectory('./landing/src');

// Resultados
const sortedEndpoints = Array.from(apiEndpoints).sort();

console.log(`📊 Total de APIs encontradas: ${sortedEndpoints.length}\n`);
console.log('═'.repeat(80));

// Agrupar por módulo
const byModule = {};
sortedEndpoints.forEach(endpoint => {
  const parts = endpoint.split('/api/')[1]?.split('/');
  const module = parts?.[0] || 'root';
  
  if (!byModule[module]) {
    byModule[module] = [];
  }
  byModule[module].push(endpoint);
});

// Mostrar por módulo
Object.keys(byModule).sort().forEach(module => {
  console.log(`\n📦 ${module.toUpperCase()} (${byModule[module].length} endpoints)`);
  console.log('─'.repeat(80));
  byModule[module].forEach(endpoint => {
    const files = filesByEndpoint.get(endpoint) || [];
    const uniqueFiles = [...new Set(files.map(f => f.split('/').slice(-2).join('/')))];
    console.log(`  ${endpoint}`);
    if (uniqueFiles.length <= 3) {
      uniqueFiles.forEach(f => console.log(`    └─ ${f}`));
    } else {
      console.log(`    └─ Usado en ${uniqueFiles.length} archivos`);
    }
  });
});

console.log('\n' + '═'.repeat(80));
console.log(`\n✅ Escaneo completo: ${sortedEndpoints.length} APIs únicas encontradas\n`);

// Guardar a archivo
import { writeFileSync } from 'fs';
const report = {
  total: sortedEndpoints.length,
  timestamp: new Date().toISOString(),
  endpoints: sortedEndpoints,
  byModule,
  filesByEndpoint: Object.fromEntries(
    Array.from(filesByEndpoint.entries()).map(([k, v]) => [k, [...new Set(v)]])
  )
};

writeFileSync('./API_SCAN_REPORT.json', JSON.stringify(report, null, 2));
console.log('💾 Reporte guardado en: API_SCAN_REPORT.json\n');
