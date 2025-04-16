export function secToFormat(sec: number) {
  if (sec === 0) {
    return `立即`;
  }
  const usedTime = sec * 1000;
  const days = Math.floor(usedTime / (24 * 3600 * 1000));
  const leavel = usedTime % (24 * 3600 * 1000);
  const hours = Math.floor(leavel / (3600 * 1000));
  const leavel2 = leavel % (3600 * 1000);
  const minutes = Math.floor(leavel2 / (60 * 1000));
  if (days > 0) {
    return `${days}天${hours}时`;
  }
  if (days === 0 && hours > 0) {
    return `${hours}时${minutes}分`;
  }
  if (days === 0 && hours === 0 && minutes > 0) {
    return `${minutes}分钟`;
  }
  return `${sec}秒`;
}

export function alertTimeToFormat(
  firstTriggerTime: string,
  recoverTime: string,
  isRecovered: boolean
) {
  let startTime = new Date();
  let endTime = new Date();
  if (isRecovered) {
    startTime = new Date(firstTriggerTime);
    endTime = new Date(recoverTime);
  } else {
    startTime = new Date(firstTriggerTime);
    endTime = new Date();
  }
  const usedTime = endTime.getTime() - startTime.getTime();
  const days = Math.floor(usedTime / (24 * 3600 * 1000));
  const leavel = usedTime % (24 * 3600 * 1000);
  const hours = Math.floor(leavel / (3600 * 1000));
  const leavel2 = leavel % (3600 * 1000);
  const minutes = Math.floor(leavel2 / (60 * 1000));
  if (days > 0) {
    return `${days}天${hours}时`;
  }
  if (days === 0 && hours > 0) {
    return `${hours}时${minutes}分`;
  }
  return `${minutes}分`;
}
