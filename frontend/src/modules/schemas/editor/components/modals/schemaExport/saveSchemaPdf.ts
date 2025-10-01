import type { jsPDF } from 'jspdf';
import type { ISchema } from '@/models';
import { collectUniqColors } from './collectUniqColors';

interface ISize {
  width: number;
  height: number;
}

function paintBackground(pdf: jsPDF, color: string) {
  const pageSize = pdf.internal.pageSize;

  pdf.setFillColor(color);
  pdf.rect(0, 0, pageSize.width, pageSize.height, 'F');
}

function replaceSourceColors(schema: ISchema, source_: string): string {
  let source = source_.replace('var(--ps-background)', schema.backgroundColor);
  const colors = collectUniqColors(schema);

  for (const [index, color] of colors.entries()) {
    source = source.replaceAll(`var(--ps-${index})`, color);
  }

  return source;
}

function buildSvgEl(schema: ISchema, source: string): SVGSVGElement {
  return new DOMParser()
    .parseFromString(replaceSourceColors(schema, source), 'image/svg+xml')
    .querySelector<SVGSVGElement>('svg')!;
}

const getSchemaSize = (svg: SVGSVGElement): ISize => ({
  width: svg.width.baseVal.value,
  height: svg.height.baseVal.value,
});

async function insertSchema(pdf: jsPDF, svgEl: SVGSVGElement): Promise<void> {
  const sourceSize = getSchemaSize(svgEl);
  const aspectRatio = sourceSize.width / sourceSize.height;

  const pageSize = pdf.internal.pageSize;
  let width = pageSize.width;
  let height = pageSize.width / aspectRatio;

  if (height > pageSize.height) {
    height = pageSize.height;
    width = pageSize.height * aspectRatio;
  }

  const x = (pageSize.width - width) / 2;
  const y = (pageSize.height - height) / 2;

  await pdf.svg(svgEl, {
    x: x,
    y: y,
    width: width,
    height: height,
  });
}

export async function saveSchemaPdf(schema: ISchema, source: string) {
  const { default: JSPDF } = await import('jspdf');
  await import('svg2pdf.js');

  const pdf = new JSPDF({
    orientation: 'landscape',
    unit: 'mm',
    format: 'a4',
  });

  paintBackground(pdf, schema.backgroundColor);
  await insertSchema(pdf, buildSvgEl(schema, source));

  await pdf.save(`${schema.name}.pdf`, { returnPromise: true });
}
