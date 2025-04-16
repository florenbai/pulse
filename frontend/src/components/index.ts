import { App } from 'vue';
import { use } from 'echarts/core';
import { CanvasRenderer } from 'echarts/renderers';
import { BarChart, LineChart, PieChart, RadarChart } from 'echarts/charts';
import {
  GridComponent,
  TooltipComponent,
  LegendComponent,
  DataZoomComponent,
  GraphicComponent,
} from 'echarts/components';
import { Icon } from '@arco-design/web-vue';
import Chart from './chart/index.vue';
import Breadcrumb from './breadcrumb/index.vue';

// 添加iconfont
const iconFont = Icon.addFromIconFontCn({
  src: 'https://at.alicdn.com/t/c/font_4395376_zlyzz5aqyk.js?spm=a313x.manage_type_myprojects.i1.10.15d63a81Wp8j8D&file=font_4395376_zlyzz5aqyk.js',
});
// Manually introduce ECharts modules to reduce packing size

use([
  CanvasRenderer,
  BarChart,
  LineChart,
  PieChart,
  RadarChart,
  GridComponent,
  TooltipComponent,
  LegendComponent,
  DataZoomComponent,
  GraphicComponent,
]);

export default {
  install(Vue: App) {
    Vue.component('Chart', Chart);
    Vue.component('Breadcrumb', Breadcrumb);
    Vue.component('IconFont', iconFont);
  },
};
