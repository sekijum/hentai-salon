import dayjs from 'dayjs';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';

dayjs.extend(utc);
dayjs.extend(timezone);

const DEFAULT_TIMEZONE = 'Asia/Tokyo';
const DEFAULT_FORMAT = 'YYYY-MM-DD HH:mm';

const formatDate = (date: Date | string, { timezone = DEFAULT_TIMEZONE, format = DEFAULT_FORMAT } = {}) => {
  if (!date || !dayjs(date).isValid()) {
    throw new Error('Invalid date provided');
  }
  return dayjs(date).tz(timezone).format(format);
};

export default defineNuxtPlugin(nuxtApp => {
  nuxtApp.provide('dayjs', dayjs);
  nuxtApp.provide('formatDate', formatDate);
});
