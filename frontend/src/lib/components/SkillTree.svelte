<script lang="ts">
  import { Canvas, Layer, t } from 'svelte-canvas';
  import type { SkillTreeData, RenderFunc, Node, Group, Sprite, Translation } from '../types';
  import { browser } from '$app/env';

  export let clickNode: (node: Node) => void;
  export let circledNode: number | undefined;

  export let selectedJewel: number;
  export let selectedConqueror: string;
  export let seed: number;

  let skillTree: SkillTreeData;

  const drawnGroups: Record<number, Group> = {};
  const drawnNodes: Record<number, Node> = {};

  const inverseSprites: Record<string, Sprite> = {};
  const inverseSpritesActive: Record<string, Sprite> = {};

  const inverseTranslations: Record<string, Translation> = {};

  const indexHandlers: Record<string, number> = {
    negate: -1,
    times_twenty: 1 / 20,
    canonical_stat: 1,
    per_minute_to_per_second: 60,
    milliseconds_to_seconds: 1000,
    display_indexable_support: 1,
    divide_by_one_hundred: 100,
    milliseconds_to_seconds_2dp_if_required: 1000,
    deciseconds_to_seconds: 10,
    old_leech_percent: 1,
    old_leech_permyriad: 10000,
    times_one_point_five: 1 / 1.5,
    '30%_of_value': 100 / 30,
    divide_by_one_thousand: 1000,
    divide_by_twelve: 12,
    divide_by_six: 6,
    per_minute_to_per_second_2dp_if_required: 60,
    '60%_of_value': 100 / 60,
    double: 1 / 2,
    negate_and_double: 1 / -2,
    multiply_by_four: 1 / 4,
    per_minute_to_per_second_0dp: 60,
    milliseconds_to_seconds_0dp: 1000,
    mod_value_to_item_class: 1,
    milliseconds_to_seconds_2dp: 1000,
    multiplicative_damage_modifier: 1,
    divide_by_one_hundred_2dp: 100,
    per_minute_to_per_second_1dp: 60,
    divide_by_one_hundred_2dp_if_required: 100,
    divide_by_ten_1dp_if_required: 10,
    milliseconds_to_seconds_1dp: 1000,
    divide_by_fifty: 50,
    per_minute_to_per_second_2dp: 60,
    divide_by_ten_0dp: 10,
    divide_by_one_hundred_and_negate: -100,
    tree_expansion_jewel_passive: 1,
    passive_hash: 1,
    divide_by_ten_1dp: 10,
    affliction_reward_type: 1,
    divide_by_five: 5,
    metamorphosis_reward_description: 1,
    divide_by_two_0dp: 2,
    divide_by_fifteen_0dp: 15,
    divide_by_three: 3,
    divide_by_twenty_then_double_0dp: 10,
    divide_by_four: 4
  };

  if (browser) {
    skillTree = JSON.parse(window['SkillTree']);

    Object.keys(skillTree.groups).forEach((groupId) => {
      const group = skillTree.groups[groupId];
      group.nodes.forEach((nodeId) => {
        const node = skillTree.nodes[nodeId];

        // Do not care about proxy passives
        if (node.isProxy) {
          return;
        }

        // Do not care about class starting nodes
        if ('classStartIndex' in node) {
          return;
        }

        // Do not care about cluster jewels
        if (node.expansionJewel) {
          if (node.expansionJewel.parent) {
            return;
          }
        }

        // Do not care about blighted nodes
        if (node.isBlighted) {
          return;
        }

        // Do not care about ascendancies
        if (node.ascendancyName) {
          return;
        }

        drawnGroups[groupId] = group;
        drawnNodes[nodeId] = node;
      });
    });

    skillTree.skillSprites.keystoneInactive.forEach((k) =>
      Object.keys(k.coords).forEach((c) => (inverseSprites[c] = k))
    );
    skillTree.skillSprites.notableInactive.forEach((k) =>
      Object.keys(k.coords).forEach((c) => (inverseSprites[c] = k))
    );
    skillTree.skillSprites.normalInactive.forEach((k) => Object.keys(k.coords).forEach((c) => (inverseSprites[c] = k)));
    skillTree.skillSprites.masteryInactive.forEach((k) =>
      Object.keys(k.coords).forEach((c) => (inverseSprites[c] = k))
    );

    skillTree.skillSprites.keystoneActive.forEach((k) =>
      Object.keys(k.coords).forEach((c) => (inverseSpritesActive[c] = k))
    );
    skillTree.skillSprites.notableActive.forEach((k) =>
      Object.keys(k.coords).forEach((c) => (inverseSpritesActive[c] = k))
    );
    skillTree.skillSprites.normalActive.forEach((k) =>
      Object.keys(k.coords).forEach((c) => (inverseSpritesActive[c] = k))
    );
    skillTree.skillSprites.masteryInactive.forEach((k) =>
      Object.keys(k.coords).forEach((c) => (inverseSpritesActive[c] = k))
    );

    const translations: Translation[] = JSON.parse(PassiveTranslations);

    translations.forEach((t) => {
      t.ids.forEach((id) => {
        inverseTranslations[id] = t;
      });
    });
  }

  const startGroups = [427, 320, 226, 227, 323, 422, 329];

  const titleFont = '25px Roboto Mono';
  const statsFont = '17px Roboto Mono';

  let scaling = 10;

  let offsetX = 0;
  let offsetY = 0;

  $: jewelRadius = 1800 / scaling;
  const included = new Set();

  type Point = {
    x: number;
    y: number;
  };

  const toCanvasCoords = (x: number, y: number): Point => {
    return {
      x: (Math.abs(skillTree.min_x) + x + offsetX) / scaling,
      y: (Math.abs(skillTree.min_y) + y + offsetY) / scaling
    };
  };

  const rotateAroundPoint = (center: Point, target: Point, angle: number): Point => {
    const radians = (Math.PI / 180) * angle;
    const cos = Math.cos(radians);
    const sin = Math.sin(radians);
    const nx = cos * (target.x - center.x) + sin * (target.y - center.y) + center.x;
    const ny = cos * (target.y - center.y) - sin * (target.x - center.x) + center.y;
    return {
      x: nx,
      y: ny
    };
  };

  // const treeImg = new Image();
  // treeImg.src = assets + '/poe_tree.png';

  const orbit16Angles = [0, 30, 45, 60, 90, 120, 135, 150, 180, 210, 225, 240, 270, 300, 315, 330];
  const orbit40Angles = [
    0, 10, 20, 30, 40, 45, 50, 60, 70, 80, 90, 100, 110, 120, 130, 135, 140, 150, 160, 170, 180, 190, 200, 210, 220,
    225, 230, 240, 250, 260, 270, 280, 290, 300, 310, 315, 320, 330, 340, 350
  ];

  const orbitAngleAt = (orbit: number, index: number): number => {
    const nodesInOrbit = skillTree.constants.skillsPerOrbit[orbit];
    if (nodesInOrbit == 16) {
      return orbit16Angles[orbit16Angles.length - index] || 0;
    } else if (nodesInOrbit == 40) {
      return orbit40Angles[orbit40Angles.length - index] || 0;
    } else {
      return 360 - (360 / nodesInOrbit) * index;
    }
  };

  const calculateNodePos = (node: Node): Point => {
    const targetGroup = skillTree.groups[node.group];
    const targetAngle = orbitAngleAt(node.orbit, node.orbitIndex);

    const targetGroupPos = toCanvasCoords(targetGroup.x, targetGroup.y);
    const targetNodePos = toCanvasCoords(targetGroup.x, targetGroup.y - skillTree.constants.orbitRadii[node.orbit]);
    return rotateAroundPoint(targetGroupPos, targetNodePos, targetAngle);
  };

  const assetCache: Record<string, HTMLImageElement> = {};
  const getAsset = (name: string): HTMLImageElement => {
    if (name in assetCache) {
      return assetCache[name];
    }

    assetCache[name] = new Image();
    assetCache[name].src = skillTree.assets[name]['0.3835'];

    return assetCache[name];
  };

  const distance = (p1: Point, p2: Point): number => {
    return Math.sqrt(Math.pow(p1.x - p2.x, 2) + Math.pow(p1.y - p2.y, 2));
  };

  $: {
    if (browser && included.size == 0) {
      skillTree.jewelSlots.forEach((jewelId) => {
        if (jewelId != 31683) {
          return;
        }

        const jewel = skillTree.nodes[jewelId];

        // Do not care about cluster jewels
        if (jewel.expansionJewel) {
          if (jewel.expansionJewel.parent) {
            return;
          }
        }

        const jRotatedPos = calculateNodePos(jewel);

        Object.keys(skillTree.nodes).forEach((nodeId) => {
          const node = skillTree.nodes[nodeId];
          const group = skillTree.groups[node.group];
          if (!group) {
            return;
          }

          const rotatedPos = calculateNodePos(node);
          if (distance(rotatedPos, jRotatedPos) < jewelRadius) {
            included.add(nodeId);
          }
        });
      });
    }
  }

  const drawScaling = 2.6;
  const drawScaledCenter = (context: CanvasRenderingContext2D, asset: HTMLImageElement, pos: Point) => {
    const newWidth = (asset.width / scaling) * drawScaling;
    const newHeight = (asset.height / scaling) * drawScaling;

    const topLeftX = pos.x - newWidth / 2;
    const topLeftY = pos.y - newHeight / 2;

    context.drawImage(asset, topLeftX, topLeftY, newWidth, newHeight);
  };

  const drawMirror = (context: CanvasRenderingContext2D, asset: HTMLImageElement, pos: Point) => {
    const newWidth = (asset.width / scaling) * drawScaling;
    const newHeight = (asset.height / scaling) * drawScaling;

    const topLeftX = pos.x - newWidth / 2;
    const topLeftY = pos.y - newHeight / 2;

    const finalY = topLeftY - newHeight / 2;

    context.drawImage(asset, topLeftX, finalY, newWidth, newHeight);
    context.save();

    context.translate(topLeftX, topLeftY);
    context.rotate(Math.PI);

    context.drawImage(asset, -newWidth, -(newHeight / 2), newWidth, -newHeight);

    context.restore();
  };

  const spriteCache: Record<string, HTMLImageElement> = {};
  const spriteCacheActive: Record<string, HTMLImageElement> = {};
  const drawSprite = (context: CanvasRenderingContext2D, path: string, pos: Point, active = false) => {
    let sprite = active ? inverseSpritesActive[path] : inverseSprites[path];

    if (!sprite && active) {
      sprite = inverseSprites[path];
    }

    const spriteSheetUrl = sprite.filename;

    if (!(spriteSheetUrl in (active ? spriteCacheActive : spriteCache))) {
      (active ? spriteCacheActive : spriteCache)[spriteSheetUrl] = new Image();
      (active ? spriteCacheActive : spriteCache)[spriteSheetUrl].src = spriteSheetUrl;
    }

    const self = sprite.coords[path];

    const newWidth = (self.w / scaling) * drawScaling;
    const newHeight = (self.h / scaling) * drawScaling;

    const topLeftX = pos.x - newWidth / 2;
    const topLeftY = pos.y - newHeight / 2;

    context.drawImage(
      (active ? spriteCacheActive : spriteCache)[spriteSheetUrl],
      self.x,
      self.y,
      self.w,
      self.h,
      topLeftX,
      topLeftY,
      newWidth,
      newHeight
    );
  };

  const wrapText = (text: string, context: CanvasRenderingContext2D, width: number): string[] => {
    const result = [];

    let currentWord = '';
    text.split(' ').forEach((word) => {
      if (context.measureText(currentWord + word).width < width) {
        currentWord += ' ' + word;
      } else {
        result.push(currentWord.trim());
        currentWord = word;
      }
    });
    61834;

    if (currentWord.length > 0) {
      result.push(currentWord.trim());
    }

    return result;
  };

  const formatStats = (translation: Translation, stat: number): string | undefined => {
    let selectedTranslation = -1;

    for (let i = 0; i < translation.English.length; i++) {
      const t = translation.English[i];

      let matches = true;
      if (t.condition.length > 0) {
        const first = t.condition[0];
        for (let condition of Object.keys(first)) {
          if (condition == 'min' && stat < first['min']) {
            matches = false;
            break;
          }

          if (condition == 'max' && stat > first['max']) {
            matches = false;
            break;
          }
        }
      }

      if (matches) {
        selectedTranslation = i;
        break;
      }
    }

    if (selectedTranslation == -1) {
      return undefined;
    }

    const datum = translation.English[selectedTranslation];

    let finalStat = stat;

    if (datum.index_handlers.length > 0) {
      datum.index_handlers[0].forEach((handler) => {
        finalStat = finalStat / (indexHandlers[handler] || 1);
      });
    }

    return datum.string.replace(`{0}`, datum.format[0].replace('#', finalStat.toString()));
  };

  let mousePos: Point = {
    x: Number.MIN_VALUE,
    y: Number.MIN_VALUE
  };

  let cursor = 'unset';

  let hoveredNode: Node | undefined;
  $: render = (({ context, width, height }) => {
    const start = window.performance.now();

    context.clearRect(0, 0, width, height);

    context.fillStyle = '#080c11';
    context.fillRect(0, 0, width, height);

    // if (!down) {
    //   const imgScale = 4.74;
    //   context.drawImage(treeImg, (offsetX - 12825) / scaling, (offsetY - 1625) / scaling, (treeImg.width * imgScale) / scaling, (treeImg.height * imgScale) / scaling);
    // }

    const connected = {};
    Object.keys(drawnGroups).forEach((groupId) => {
      context.strokeStyle = `hsl(${$t / 20}, 100%, 50%)`;

      if (groupId != 329) {
        // return;
      }

      const group = skillTree.groups[groupId];
      const groupPos = toCanvasCoords(group.x, group.y);

      const maxOrbit = Math.max(...group.orbits);
      if (startGroups.indexOf(parseInt(groupId)) >= 0) {
        // Do not draw starter nodes
      } else if (maxOrbit == 1) {
        drawScaledCenter(context, getAsset('PSGroupBackground1'), groupPos);
      } else if (maxOrbit == 2) {
        drawScaledCenter(context, getAsset('PSGroupBackground2'), groupPos);
      } else if (maxOrbit == 3 || group.orbits.length > 1) {
        drawMirror(context, getAsset('PSGroupBackground3'), groupPos);
      }
    });

    Object.keys(drawnNodes).forEach((nodeId) => {
      const node = skillTree.nodes[nodeId];
      const angle = orbitAngleAt(node.orbit, node.orbitIndex);
      const rotatedPos = calculateNodePos(node);

      node.out?.forEach((o) => {
        if (!drawnNodes[parseInt(o)]) {
          return;
        }

        const min = Math.min(parseInt(o), parseInt(nodeId));
        const max = Math.max(parseInt(o), parseInt(nodeId));
        const joined = min + ':' + max;

        if (joined in connected) {
          return;
        }
        connected[joined] = true;

        const targetNode = skillTree.nodes[parseInt(o)];

        // Do not draw connections to mastery nodes
        if (targetNode.isMastery) {
          return;
        }

        const targetAngle = orbitAngleAt(targetNode.orbit, targetNode.orbitIndex);
        const targetRotatedPos = calculateNodePos(targetNode);

        context.beginPath();

        if (node.group != targetNode.group || node.orbit != targetNode.orbit) {
          context.moveTo(rotatedPos.x, rotatedPos.y);
          context.lineTo(targetRotatedPos.x, targetRotatedPos.y);
        } else {
          let a = Math.PI / 180 - (Math.PI / 180) * angle;
          let b = Math.PI / 180 - (Math.PI / 180) * targetAngle;

          a -= Math.PI / 2;
          b -= Math.PI / 2;

          const diff = Math.abs(Math.max(a, b) - Math.min(a, b));

          const finalA = diff > Math.PI ? Math.max(a, b) : Math.min(a, b);
          const finalB = diff > Math.PI ? Math.min(a, b) : Math.max(a, b);

          const group = skillTree.groups[node.group];
          const groupPos = toCanvasCoords(group.x, group.y);
          context.arc(groupPos.x, groupPos.y, skillTree.constants.orbitRadii[node.orbit] / scaling + 1, finalA, finalB);
        }

        context.lineWidth = 6 / scaling;
        context.strokeStyle = `#524518`;
        context.stroke();
      });
    });

    let circledNodePos: Point;
    if (circledNode) {
      circledNodePos = calculateNodePos(skillTree.nodes[circledNode]);
      context.strokeStyle = '#ad2b2b';
    }

    let hoveredNodeActive = false;
    let newHoverNode: Node | undefined;
    Object.keys(drawnNodes).forEach((nodeId) => {
      const node = skillTree.nodes[nodeId];
      const rotatedPos = calculateNodePos(node);
      let touchDistance = 0;

      let active = false;
      if (circledNode) {
        if (distance(rotatedPos, circledNodePos) < jewelRadius) {
          active = true;
        }
      }

      if (node.isKeystone) {
        touchDistance = 110;
        drawSprite(context, node.icon, rotatedPos, active);
        if (active) {
          drawScaledCenter(context, getAsset('KeystoneFrameAllocated'), rotatedPos);
        } else {
          drawScaledCenter(context, getAsset('KeystoneFrameUnallocated'), rotatedPos);
        }
      } else if (node.isNotable) {
        touchDistance = 70;
        drawSprite(context, node.icon, rotatedPos, active);
        if (active) {
          drawScaledCenter(context, getAsset('NotableFrameAllocated'), rotatedPos);
        } else {
          drawScaledCenter(context, getAsset('NotableFrameUnallocated'), rotatedPos);
        }
      } else if (node.isJewelSocket) {
        touchDistance = 70;
        if (node.expansionJewel) {
          if (active) {
            drawScaledCenter(context, getAsset('JewelSocketAltNormal'), rotatedPos);
          } else {
            drawScaledCenter(context, getAsset('JewelSocketAltNormal'), rotatedPos);
          }
        } else {
          if (active) {
            drawScaledCenter(context, getAsset('JewelFrameAllocated'), rotatedPos);
          } else {
            drawScaledCenter(context, getAsset('JewelFrameUnallocated'), rotatedPos);
          }
        }
      } else if (node.isMastery) {
        drawSprite(context, node.inactiveIcon, rotatedPos, active);
      } else {
        touchDistance = 50;
        drawSprite(context, node.icon, rotatedPos, active);
        if (active) {
          drawScaledCenter(context, getAsset('PSSkillFrameActive'), rotatedPos);
        } else {
          drawScaledCenter(context, getAsset('PSSkillFrame'), rotatedPos);
        }
      }

      if (distance(rotatedPos, mousePos) < touchDistance / scaling) {
        newHoverNode = node;
        hoveredNodeActive = active;
      }
    });

    hoveredNode = newHoverNode;

    // Object.keys(drawnGroups).forEach(groupId => {
    //   context.strokeStyle = `hsl(${$t / 20}, 100%, 50%)`;
    //
    //   const group = skillTree.groups[groupId];
    //   const groupPos = toCanvasCoords(group.x, group.y);
    //
    //   group.orbits.forEach(o => {
    //     context.beginPath();
    //     context.arc(groupPos.x, groupPos.y, skillTree.constants.orbitRadii[o] / scaling, 0, Math.PI * 2);
    //     context.stroke();
    //   });
    // });

    if (circledNode) {
      context.strokeStyle = '#ad2b2b';
      context.beginPath();
      context.arc(circledNodePos.x, circledNodePos.y, jewelRadius, 0, Math.PI * 2);
      context.stroke();
    }

    if (hoveredNode) {
      let nodeName = hoveredNode.name;
      let nodeStats: { text: string; special: boolean }[] = (hoveredNode.stats || []).map((s) => ({
        text: s,
        special: false
      }));

      if (!hoveredNode.isJewelSocket && hoveredNodeActive) {
        if (hoveredNode.skill && seed && selectedJewel && selectedConqueror) {
          const result = Calculate(TreeToPassive[hoveredNode.skill], seed, selectedJewel, selectedConqueror);

          if (result) {
            if ('alternatePassiveSkill' in result) {
              const alt = GetAlternatePassiveSkillByIndex(result['alternatePassiveSkill']);

              nodeStats = [];
              nodeName = alt.name;

              if ('statsKeys' in alt) {
                alt['statsKeys'].forEach((statId, i) => {
                  const stat = GetStatByIndex(statId);
                  const translation = inverseTranslations[stat.id] || '';
                  if (translation) {
                    nodeStats.push({
                      text: formatStats(translation, result.statRolls[i]) || stat.id,
                      special: true
                    });
                  }
                });
              }
            }

            if (
              'alternatePassiveAdditionInformations' in result &&
              result['alternatePassiveAdditionInformations'].length > 0
            ) {
              result['alternatePassiveAdditionInformations'].forEach((info) => {
                const addition = GetAlternatePassiveAdditionByIndex(info.alternatePassiveSkillAddition);

                if ('statsKeys' in addition) {
                  addition['statsKeys'].forEach((statId, i) => {
                    const stat = GetStatByIndex(statId);
                    const translation = inverseTranslations[stat.id] || '';
                    if (translation) {
                      nodeStats.push({
                        text: formatStats(translation, info.statRolls[i]) || stat.id,
                        special: true
                      });
                    }
                  });
                }
              });
            }
          }
        }
      }

      context.font = titleFont;
      const textMetrics = context.measureText(nodeName);

      const width = Math.max(textMetrics.width + 50, 600);

      context.font = statsFont;

      const allLines: {
        text: string;
        offset: number;
        special: boolean;
      }[] = [];

      const padding = 30;

      let offset = 85;

      if (nodeStats && nodeStats.length > 0) {
        nodeStats.forEach((stat) => {
          if (allLines.length > 0) {
            offset += 5;
          }

          stat.text.split('\n').forEach((line) => {
            if (allLines.length > 0) {
              offset += 10;
            }

            const lines = wrapText(line, context, width - padding);
            lines.forEach((l) => {
              allLines.push({
                text: l,
                offset,
                special: stat.special
              });
              offset += 20;
            });
          });
        });
      } else if (hoveredNode.isJewelSocket) {
        allLines.push({
          text: 'Click to select this socket',
          offset,
          special: true
        });

        offset += 20;
      }

      const titleHeight = 55;

      context.fillStyle = 'rgba(75,63,24,0.9)';
      context.fillRect(mousePos.x, mousePos.y, width, titleHeight);

      context.fillStyle = '#ffffff';
      context.font = titleFont;
      context.textAlign = 'center';
      context.fillText(nodeName, mousePos.x + width / 2, mousePos.y + 35);

      context.fillStyle = 'rgba(0,0,0,0.8)';
      context.fillRect(mousePos.x, mousePos.y + titleHeight, width, offset - titleHeight);

      context.font = statsFont;
      context.textAlign = 'left';
      allLines.forEach((l) => {
        if (l.special) {
          context.fillStyle = '#8cf34c';
        } else {
          context.fillStyle = '#ffffff';
        }

        context.fillText(l.text, mousePos.x + padding / 2, mousePos.y + l.offset);
      });
    }

    if (hoveredNode && hoveredNode.isJewelSocket) {
      cursor = 'pointer';
    } else {
      cursor = 'unset';
    }

    context.fillStyle = '#ffffff';
    context.textAlign = 'right';
    context.font = '12px Roboto Mono';

    const end = window.performance.now();

    context.fillText(`${(end - start).toFixed(1)}ms`, width - 5, 17);
  }) as RenderFunc;

  let downX = 0;
  let downY = 0;

  let startX = 0;
  let startY = 0;

  let down = false;
  const mouseDown = (event: MouseEvent) => {
    down = true;
    downX = event.offsetX;
    downY = event.offsetY;
    startX = offsetX;
    startY = offsetY;

    mousePos = {
      x: event.offsetX,
      y: event.offsetY
    };

    if (hoveredNode) {
      clickNode(hoveredNode);
    }
  };

  const mouseUp = (event: PointerEvent) => {
    if (event.type === 'pointerup') {
      down = false;
    }

    mousePos = {
      x: event.offsetX,
      y: event.offsetY
    };
  };

  const mouseMove = (event: MouseEvent) => {
    if (down) {
      offsetX = startX - (downX - event.offsetX) * scaling;
      offsetY = startY - (downY - event.offsetY) * scaling;
    }

    mousePos = {
      x: event.offsetX,
      y: event.offsetY
    };
  };

  const onScroll = (event: WheelEvent) => {
    if (event.deltaY > 0) {
      if (scaling < 30) {
        offsetX += event.offsetX;
        offsetY += event.offsetY;
      }
    } else {
      if (scaling > 3) {
        offsetX -= event.offsetX;
        offsetY -= event.offsetY;
      }
    }

    scaling = Math.min(30, Math.max(3, scaling + event.deltaY / 100));

    event.preventDefault();
    event.stopPropagation();
    event.stopImmediatePropagation();
  };

  let width = 0;
  let height = 0;
  const resize = () => {
    width = window.innerWidth;
    height = window.innerHeight;
  };

  let initialized = false;
  $: {
    if (browser) {
      if (!initialized) {
        initialized = true;
        offsetX = skillTree.min_x + (window.innerWidth / 2) * scaling;
        offsetY = skillTree.min_y + (window.innerHeight / 2) * scaling;
        resize();
      }
    }
  }
</script>

<svelte:window on:pointerup={mouseUp} on:pointermove={mouseMove} on:resize={resize} />

{#if browser && width && height}
  <div on:resize={resize} style="touch-action: none; cursor: {cursor}">
    <Canvas {width} {height} on:pointerdown={mouseDown} on:wheel={onScroll}>
      <Layer {render} />
    </Canvas>
    <slot />
  </div>
{/if}
