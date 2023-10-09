import * as esbuild from "esbuild";
import postCssPlugin from "esbuild-style-plugin";
import autoprefixer from "autoprefixer";
import plugin from "tailwindcss";

const args = process.argv.slice(2);
const watch = args.includes("--watch");

const watchPlugin = {
  name: "watch-plugin",
  setup(build) {
    build.onStart(() => {
      console.log(`Build starting: ${new Date(Date.now()).toLocaleString()}`);
    });
    build.onEnd((result) => {
      if (result.errors.length > 0) {
        console.log(
          `Build finished, with errors: ${new Date(
            Date.now()
          ).toLocaleString()}`
        );
      } else {
        console.log(
          `Build finished successfully: ${new Date(
            Date.now()
          ).toLocaleString()}`
        );
      }
    });
  },
};

let ctx = await esbuild.context({
  entryPoints: ["frontend/App.tsx", "frontend/style.css"],
  outdir: "public/assets",
  bundle: true,
  minify: true,
  plugins: [
    watchPlugin,
    postCssPlugin({
      postcss: {
        plugins: [plugin, autoprefixer],
      },
    }),
  ],
});

if (watch) {
  await ctx.watch();
  console.log("watching...");
} else {
  ctx.rebuild();
  ctx.dispose();
}

// console.log(ctx)
// await ctx.watch();

// build({
//   entryPoints: ["frontend/App.tsx", "frontend/style.css"],
//   outdir: "public/assets",
//   bundle: true,
//   // watch: true,
//   minify: true,
//   plugins: [
//     postCssPlugin({
//       postcss: {
//         plugins: [plugin, autoprefixer],
//       },
//     }),
//   ],
// })
//   .then(() => console.log("⚡ Build complete! ⚡"))
//   .catch(() => {
//     console.error(`Build error: ${error}`);
//     process.exit(1);
//   });
