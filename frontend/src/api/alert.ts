import axios from 'axios';

export interface AlertTemplateRecode {
  id: number;
  templateName: string;
}

export function queryAllTemplate() {
  return axios.get(`/api/v1/alert/template/all`);
}
